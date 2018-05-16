package main

import (
	"log"
	"net/http"

	"github.com/akardc/tcabaagno/server/internal/routes"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	routes.Init(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
