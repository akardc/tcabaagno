package websockets

import (
	"encoding/json"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/satori/go.uuid"
)

const (
	StartGameMessage            string = "start-game"
	RegistrationResponseMessage string = "joined-game"
	GameUpdatedMessage          string = "game-updated"
	SubmitQuestionsMessage      string = "submit-questions"
)

type SubmitQuestionsMessagePayload struct {
	Who   string `json:"who"`
	What  string `json:"what"`
	When  string `json:"when"`
	Where string `json:"where"`
	Why   string `json:"why"`
}

type RegistrationResponseMessagePayload struct {
	ID     string `json:"id"`
	GameID string `json:"gameId"`
}

type GameUpdatedMessagePayload struct {
	ID         string `json:"id"`
	NumPlayers int    `json:"numPlayers"`
}

type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// GameController maintains the set of active clients and broadcasts messages to the
// clients.
type GameController struct {
	ID string

	// Registered clients.
	// clients map[*Client]bool
	clients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewGameController() *GameController {
	name := randomdata.Adjective() + "-" + randomdata.Noun()
	return &GameController{
		ID:         name,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
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
		case message := <-g.broadcast:
			g.handleReceivedMessage(message)
		}
	}
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
			ID:         g.ID,
			NumPlayers: len(g.clients),
		},
	})
}

func (g *GameController) handleReceivedMessage(messageData []byte) {
	var message Message
	if err := json.Unmarshal(messageData, &message); err != nil {
		log.Printf("Error decoding message: %s", messageData)
		return
	}
	log.Printf("Routing Incoming Message: %s", message.Type)
	switch message.Type {
	case StartGameMessage:
		g.startGame()
	default:
		log.Printf("No handler for message \"%s\", discarding", message)
	}
}

func (g *GameController) startGame() {
	g.sendMessageToAllClients(Message{
		Type: StartGameMessage,
	})
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
