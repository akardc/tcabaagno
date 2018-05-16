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
		games: map[string]*websockets.Hub{},
	}
}

type GameOrch struct {
	hub   *websockets.Hub
	games map[string]*websockets.Hub
}

func (o *GameOrch) CreateNewGame() (NewGameOutput, error) {
	gameName := randomdata.Adjective() + "-" + randomdata.Noun()
	o.games[gameName] = websockets.NewHub()
	go o.games[gameName].Run()
	log.Printf("Started new game with name %s", gameName)
	return NewGameOutput{
		Name: gameName,
	}, nil
}

func (o *GameOrch) Join(gameName string, w http.ResponseWriter, r *http.Request) error {
	game, ok := o.games[gameName]
	if !ok {
		return fmt.Errorf("The game %s was not found", gameName)
	}
	websockets.ServeWS(game, w, r)
	return nil
}

func (o *GameOrch) ListActiveGames() {

}

type NewGameOutput struct {
	Name string
}
