package main

import (
	"log"
	"net/http"
	"github.com/spiron09/burnervpn/server/handlers"
)


func main() {
	http.HandleFunc("/regions", handlers.HandleRegions)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
