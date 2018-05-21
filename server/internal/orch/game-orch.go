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
	log.Printf("Started new game with ID %s", game.ID)
	return &GameOutput{
		ID: game.ID,
	}, nil
}

func (o *GameOrch) GetActiveGames() ([]*GameOutput, error) {
	var output []*GameOutput

	for _, game := range o.games {
		output = append(output, &GameOutput{
			ID: game.ID,
			Status: game.CurrentGameStep,
		})
	}

	return output, nil
}

func (o *GameOrch) Join(gameID string, w http.ResponseWriter, r *http.Request) error {
	game, ok := o.games[gameID]
	if !ok {
		return fmt.Errorf("The game %s was not found", gameID)
	}
	if err := game.Join(w, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	return nil
}

type GameController struct {
	ID             string
	GameController *websockets.GameController
}

type GameOutput struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
