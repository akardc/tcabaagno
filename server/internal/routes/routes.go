package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Init(router chi.Router) error {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"content-type"},
	})
	router.Use(corsHandler.Handler)
	router.Mount(NewGameRoutes("/games").Router())

	return nil
}
