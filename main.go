package main

import (
	"net/http"

	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/db"
	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/models"
	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.User{}, models.Task{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
