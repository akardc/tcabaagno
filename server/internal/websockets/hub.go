package websockets

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/satori/go.uuid"
)

const (
	StartGameMessage       string = "start-game"
	PlayerJoinedMessage    string = "player-joined"
	SubmitQuestionsMessage string = "submit-questions"
)

type SubmitQuestionsMessagePayload struct {
	Who   string `json:"who"`
	What  string `json:"what"`
	When  string `json:"when"`
	Where string `json:"where"`
	Why   string `json:"why"`
}

type PlayerJoinedMessagePayload struct {
	ID string `json:"id"`
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
		Type: PlayerJoinedMessage,
		Payload: PlayerJoinedMessagePayload{
			ID: client.ID,
		},
	})
}

func (g *GameController) handleReceivedMessage(message []byte) {
	msg := strings.Trim(string(message), "\"")
	log.Printf("Routing Incoming Message: %s", msg)
	switch msg {
	case StartGameMessage:
		g.sendMessageToAllClients(msg)
	default:
		log.Printf("No handler for message \"%s\", discarding", msg)
	}
}

func (g *GameController) sendMessageToSingleClient(client *Client, message Message) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message to json: %e", err)
		return
	}
	client.send <- []byte(messageJSON)
}

func (g *GameController) sendMessageToAllClients(messageType string) {
	log.Printf("Sending message to all clients: %s", string(messageType))
	var message Message
	message = Message{
		Type: messageType,
	}
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
