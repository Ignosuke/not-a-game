package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ignosuke/not_a_game-api-rest/common"
	"github.com/Ignosuke/not_a_game-api-rest/models"
	"github.com/gorilla/mux"
)

func Save(writer http.ResponseWriter, request *http.Request) { //Create||Update
	user := models.Users{}

	db := common.GetConnection()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Create(&user).Error

	if err != nil {
		log.Fatal(err)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)
	common.SendResponse(writer, http.StatusOK, json)
}

func GetAll(writer http.ResponseWriter, request *http.Request) { //Read all
	user := []models.Users{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&user)

	json, _ := json.Marshal(user)

	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) { //Read one
	user := models.Users{}

	id := mux.Vars(request)["id"]
	db := common.GetConnection()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}

func Delete(writer http.ResponseWriter, request *http.Request) { //Delete
	user := models.Users{}
	id := mux.Vars(request)["id"]
	db := common.GetConnection()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		db.Delete(user)
		common.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}
