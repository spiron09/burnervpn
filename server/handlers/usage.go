package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type UsageResponse struct {
	SessionID         string  `json:"session_id"`
	Cost              float64 `json:"cost"`
	DurationInSeconds float64 `json:"duration"`
}

func HandleUsage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sessionID := vars["id"]

	session := sessionStore.GetSession(sessionID)
	if session == nil {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	duration := time.Since(session.CreatedAt)
	hours := duration.Hours()
	cost := hours * 0.015

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UsageResponse{
		SessionID:         sessionID,
		Cost:              cost,
		DurationInSeconds: duration.Seconds(),
	})
}
