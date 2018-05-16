package orch

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pallinder/go-randomdata"
	"github.com/akardc/tcabaagno/server/internal/websockets"
)

func NewGameOrch() *GameOrch {
	return &GameOrch{
		games: map[string]*GameController{},
	}
}

type GameOrch struct {
	games map[string]*GameController
}

func (o *GameOrch) CreateNewGame() (*GameOutput, error) {
	game := NewGameController()
	o.games[game.Name] = game
	log.Printf("Started new game with name %s", game.Name)
	return &GameOutput{
		Name: game.Name,
	}, nil
}

func (o *GameOrch) GetActiveGames() ([]*GameOutput, error) {
	var output []*GameOutput

	for _, game := range o.games {
		output = append(output, &GameOutput{
			Name: game.Name,
		})
	}

	return output, nil
}

func (o *GameOrch) Join(gameName string, w http.ResponseWriter, r *http.Request) error {
	game, ok := o.games[gameName]
	if !ok {
		return fmt.Errorf("The game %s was not found", gameName)
	}
	websockets.ServeWS(game.Hub, w, r)
	return nil
}

type GameController struct {
	Name string
	Hub  *websockets.Hub
}

func NewGameController() *GameController {
	name := randomdata.Adjective() + "-" + randomdata.Noun()
	hub := websockets.NewHub()
	go hub.Run()
	return &GameController{
		Name: name,
		Hub:  hub,
	}
}

type GameOutput struct {
	Name string
}
