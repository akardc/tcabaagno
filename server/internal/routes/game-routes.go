package routes

import (
	"net/http"

	"github.com/akardc/tcabaagno/server/internal/orch"
	"github.com/go-chi/chi"
)

type GameRoutes struct {
	mountPath string
	gameOrch  *orch.GameOrch
}

func NewGameRoutes(mountPath string) *GameRoutes {
	return &GameRoutes{
		mountPath: mountPath,
		gameOrch:  orch.NewGameOrch(),
	}
}

func (gr *GameRoutes) Router() (string, chi.Router) {
	router := chi.NewRouter()

	router.Get("/status", RequestHandler(func() (interface{}, error) {
		return struct {
			Status string
		}{
			Status: "ok",
		}, nil
	}))

	router.Post("/", RequestHandler(func() (interface{}, error) {
		return gr.gameOrch.CreateNewGame()
	}))

	router.Get("/", RequestHandler(func() (interface{}, error) {
		return gr.gameOrch.GetActiveGames()
	}))

	router.Get("/{gameName}/join", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gameName := chi.URLParam(r, "gameName")
		gr.gameOrch.Join(gameName, w, r)
	}))

	return gr.mountPath, router
}
