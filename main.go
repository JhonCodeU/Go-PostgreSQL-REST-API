package main

import (
	"net/http"

	"github.com/JhonCodeU/Go-PostgreSQL-REST-API-/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":8000", r)
}
