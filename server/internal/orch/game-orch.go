package orch

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akardc/tcabaagno/server/internal/websockets"
)

func NewGameOrch() *GameOrch {
	return &GameOrch{
		games: map[string]*websockets.GameController{},
	}
}

type GameOrch struct {
	games map[string]*websockets.GameController
}

func (o *GameOrch) CreateNewGame() (*GameOutput, error) {
	game := websockets.NewGameController()
	o.games[game.ID] = game
	go game.Run()
	log.Printf("Started new game with name %s", game.ID)
	return &GameOutput{
		Name: game.ID,
	}, nil
}

func (o *GameOrch) GetActiveGames() ([]*GameOutput, error) {
	var output []*GameOutput

	for _, game := range o.games {
		output = append(output, &GameOutput{
			Name: game.ID,
		})
	}

	return output, nil
}

func (o *GameOrch) Join(gameName string, w http.ResponseWriter, r *http.Request) error {
	game, ok := o.games[gameName]
	if !ok {
		return fmt.Errorf("The game %s was not found", gameName)
	}
	websockets.ServeWS(game, w, r)
	return nil
}

type GameController struct {
	Name           string
	GameController *websockets.GameController
}

type GameOutput struct {
	Name string
}
