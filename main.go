package main

import (
	"log"
	"net/http"

	"github.com/Ignosuke/not_a_game-api/common"
	"github.com/Ignosuke/not_a_game-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	common.Migrate()
	router := mux.NewRouter()
	routes.SetUserRoutes(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Server running on port 3000")
	log.Println(server.ListenAndServe())
}
