package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spiron09/burnervpn/server/config"
)

type DeleteSessionResponse struct {
	SessionID string  `json:"session_id"`
	Duration  float64 `json:"duration"`
	Cost      float64 `json:"cost"`
}

func HandleDeleteSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]

	if sessionID == "" {
		http.Error(w, "Missing session ID", http.StatusBadRequest)
		return
	}

	session := sessionStore.GetSession(sessionID)
	if session == nil {
		http.Error(w, "Sesssion not found", http.StatusNotFound)
		return
	}

	duration := time.Since(session.CreatedAt)

	client := config.LoadDOClient()
	fmt.Println("hi!!")
	_, err := client.Droplets.Delete(context.Background(), session.DropletID)
	if err != nil {
		http.Error(w, "Failed to delete droplet", http.StatusInternalServerError)
		return
	}
	fmt.Println("hi!!")
	hours := duration.Hours()
	cost := hours * 0.015

	sessionStore.DeleteSession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DeleteSessionResponse{
		SessionID: sessionID,
		Duration:  duration.Seconds(),
		Cost:      cost,
	})
}
