package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spiron09/burnervpn/server/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sessions", handlers.HandleCreateSession).Methods("POST")
	r.HandleFunc("/sessions/{id}", handlers.HandleDeleteSession).Methods("DELETE")
	r.HandleFunc("/sessions/{id}/usage", handlers.HandleUsage).Methods("GET")

	r.HandleFunc("/regions", handlers.HandleRegions).Methods("GET")
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
