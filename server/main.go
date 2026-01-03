package main

import (
	"github.com/spiron09/burnervpn/server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sessions", handlers.HandleCreateSession)
	http.HandleFunc("/regions", handlers.HandleRegions)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
