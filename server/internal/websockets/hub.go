package websockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/akardc/tcabaagno/server/internal/model"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

const (
	StartGameMessage            string = "start-game"
	RegistrationResponseMessage string = "joined-game"
	GameUpdatedMessage          string = "game-updated"
	SubmitQuestionsMessage      string = "submit-questions"
	QuestionsOutputMessage      string = "questions-output"

	AcceptingPlayersGameStep   string = "accepting-players"
	AcceptingQuestionsGameStep string = "accepting-questions"
	AcceptingAnswersGameStep   string = "accepting-answers"
	ReadingSubmissionsGameStep string = "reading-submissions"
)

type SubmitQuestionsMessagePayload struct {
	Who      string `json:"who"`
	What     string `json:"what"`
	When     string `json:"when"`
	Where    string `json:"where"`
	Why      string `json:"why"`
	PlayerID string `json:"playerId"`
}

type RegistrationResponseMessagePayload struct {
	ID     string `json:"id"`
	GameID string `json:"gameId"`
}

type GameUpdatedMessagePayload struct {
	ID          string `json:"id"`
	NumPlayers  int    `json:"numPlayers"`
	CurrentStep string `json:"currentStep"`
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type MessageInput struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type BroadcastInput struct {
	ClientID string
	Message  []byte
}

// GameController maintains the set of active clients and broadcasts messages to the
// clients.
type GameController struct {
	ID              string
	CurrentGameStep string

	QAForms map[string]model.QAForm

	// Registered clients.
	// clients map[*Client]bool
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan BroadcastInput

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewGameController() *GameController {
	name := randomdata.Adjective() + "-" + randomdata.Noun()
	return &GameController{
		ID:              name,
		CurrentGameStep: AcceptingPlayersGameStep,
		QAForms:         make(map[string]model.QAForm),
		broadcast:       make(chan BroadcastInput),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		clients:         make(map[string]*Client),
	}
}

func (g *GameController) Run() {
	for {
		select {
		case client := <-g.register:
			g.registerClient(client)
		case client := <-g.unregister:
			if _, ok := g.clients[client.ID]; ok {
				delete(g.clients, client.ID)
				close(client.send)
				g.gameUpdated()
			}
		case broadcastInput := <-g.broadcast:
			g.handleReceivedMessage(broadcastInput)
		}
	}
}

func (g *GameController) Join(w http.ResponseWriter, r *http.Request) error {
	if g.CurrentGameStep == AcceptingPlayersGameStep {
		ServeWS(g, w, r)
	} else {
		return errors.Errorf("Could not join game. It has already started")
	}
	return nil
}

func (g *GameController) registerClient(client *Client) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed to set an ID for the new client")
	}
	client.ID = id.String()
	g.clients[client.ID] = client

	g.sendMessageToSingleClient(client, Message{
		Type: RegistrationResponseMessage,
		Payload: RegistrationResponseMessagePayload{
			ID:     client.ID,
			GameID: g.ID,
		},
	})
	g.gameUpdated()
}

func (g *GameController) gameUpdated() {
	g.sendMessageToAllClients(Message{
		Type: GameUpdatedMessage,
		Payload: GameUpdatedMessagePayload{
			ID:          g.ID,
			NumPlayers:  len(g.clients),
			CurrentStep: g.CurrentGameStep,
		},
	})
}

func (g *GameController) handleReceivedMessage(broadcastInput BroadcastInput) {
	var message MessageInput
	if err := json.Unmarshal(broadcastInput.Message, &message); err != nil {
		log.Println(broadcastInput.Message)
		log.Printf("Error decoding message: %s\n\tCause: %s", broadcastInput.Message, err.Error())
		return
	}
	log.Printf("Routing Incoming Message: %s", message.Type)
	log.Printf("MessageInput: %+v", message)
	switch message.Type {
	case StartGameMessage:
		g.startGame()
	case SubmitQuestionsMessage:
		g.receiveQuestionsForm(message, broadcastInput.ClientID)
	default:
		log.Printf("No handler for message \"%s\", discarding", message)
	}
}

func (g *GameController) startGame() {
	if g.CurrentGameStep == AcceptingPlayersGameStep {
		g.CurrentGameStep = AcceptingQuestionsGameStep
		g.sendMessageToAllClients(Message{
			Type: StartGameMessage,
		})
		g.gameUpdated()
	}
}

func (g *GameController) receiveQuestionsForm(message MessageInput, submittingClientID string) {
	var questionsInput SubmitQuestionsMessagePayload
	mapstructure.Decode(message.Payload, &questionsInput)
	questions := model.Questions{
		Who:      questionsInput.Who,
		What:     questionsInput.What,
		When:     questionsInput.When,
		Where:    questionsInput.Where,
		Why:      questionsInput.Why,
		PlayerID: submittingClientID,
	}
	log.Printf("Received questions: %+v", questions)
	newForm := model.QAForm{
		Questions: questions,
	}
	if _, ok := g.QAForms[submittingClientID]; !ok {
		g.QAForms[submittingClientID] = newForm
	}
	g.advanceToAnswersIfReady()
}

func (g *GameController) advanceToAnswersIfReady() {
	if len(g.QAForms) == len(g.clients) {
		g.CurrentGameStep = AcceptingAnswersGameStep
	}
	g.gameUpdated()
}

func (g *GameController) sendQuestions() {
	forms := g.getShuffledForms(0)
	for id, form := range forms {
		g.sendMessageToSingleClient(g.clients[id], Message{
			Type: QuestionsOutputMessage,
			Payload: SubmitQuestionsMessagePayload{
				Who:      form.Questions.Who,
				What:     form.Questions.What,
				When:     form.Questions.When,
				Where:    form.Questions.Where,
				Why:      form.Questions.Why,
				PlayerID: form.Questions.PlayerID,
			},
		})
	}
}

func (g *GameController) getShuffledForms(depth int) map[string]model.QAForm {
	//
	return map[string]model.QAForm{}
}

func (g *GameController) sendMessageToSingleClient(client *Client, message Message) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message to json: %e", err)
		return
	}
	client.send <- []byte(messageJSON)
}

func (g *GameController) sendMessageToAllClients(message Message) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return
	}
	for id, client := range g.clients {
		select {
		case client.send <- []byte(messageJSON):
		default:
			close(client.send)
			delete(g.clients, id)
		}
	}
}
