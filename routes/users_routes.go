package routes

import (
	"github.com/Ignosuke/not_a_game-api-rest/controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
