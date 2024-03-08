package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/rooms", CreateRoom).Methods("POST")
	r.HandleFunc("/rooms/{id}/join", JoinRoom).Methods("POST")

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
