package main

import (
	"beintil/mongo-http/internal/adapters/api/note"
	"beintil/mongo-http/internal/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = "8080"
)

var router = mux.NewRouter()

func Handler() {
	router.HandleFunc("/notes/", note.GetAllNotes).Methods("GET")        // All Note
	router.HandleFunc("/notes/", note.CreateNote).Methods("POST")        // Create Note
	router.HandleFunc("/notes/{id}", note.UpdateNote).Methods("PATCH")  // Update Note
	router.HandleFunc("/notes/{id}", note.DeleteNote).Methods("DELETE") // Delete Note
}

func main() {
	log.Println("Start server...\n", "")

	database.ConnectDB()
	Handler()

	log.Printf("Start server on localhost:%s", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Port connection error:%s\n%s", port, err)
	}
}
