package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/db"
	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)
	id := vars["id"]
	db.DB.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}