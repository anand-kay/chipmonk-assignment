package main

import (
	"net/http"

	"github.com/anand-kay/chipmonk-assignment/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	regHandlers(r)

	http.ListenAndServe("localhost:3000", r)
}

func regHandlers(r *mux.Router) {
	r.HandleFunc("/register", handlers.RegisterHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/activeusers", handlers.ActiveUsersHandler).Methods(http.MethodGet)
}
