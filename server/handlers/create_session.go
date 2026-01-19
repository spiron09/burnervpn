package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/digitalocean/godo"
	"github.com/google/uuid"
	"github.com/spiron09/burnervpn/server/config"
	"github.com/spiron09/burnervpn/server/models"
	"github.com/spiron09/burnervpn/server/services"
	"github.com/spiron09/burnervpn/server/store"
)

type CreateSessionRequest struct {
	Region string `json:"region"`
}

type CreateSessionResponse struct {
	SessionID       string `json:"session_id"`
	WireGuardConfig string `json:"wireguard_config"`
}

var sessionStore = store.SessionStoreInit()

func HandleCreateSession(w http.ResponseWriter, r *http.Request) {
	var req CreateSessionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	serverKeys, _ := services.GenerateKeyPair()
	clientKeys, _ := services.GenerateKeyPair()

	session := &models.Session{
		ID:            uuid.New().String(),
		Region:        req.Region,
		Status:        "Creating",
		CreatedAt:     time.Now().UTC(),
		ServerKeyPair: serverKeys,
		ClientKeyPair: clientKeys,
	}

	userData := services.BuildServerConfig(serverKeys.PrivateKey, clientKeys.PublicKey)
	client := config.LoadDOClient()

	dropletReq := &godo.DropletCreateRequest{
		Name:   "burner-" + session.ID[:8],
		Region: req.Region,
		Size:   "s-1vcpu-512mb-10gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-22-04-x64",
		},
		UserData: userData,
	}

	droplet, _, err := client.Droplets.Create(context.Background(), dropletReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	provisionedDroplet, err := services.WaitForDropletProvision(client, droplet.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.DropletID = provisionedDroplet.ID
	session.DropletIP = provisionedDroplet.Networks.V4[0].IPAddress
	session.Status = "Active"
	sessionStore.SetSession(session)
	fmt.Printf("Droplet Provisioned Sucesfully with IP %s\n", session.DropletIP)

	// err = services.WaitForWireGuardReady(session.DropletIP)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	clientConfig := services.BuildClientConfig(serverKeys.PublicKey, clientKeys.PrivateKey, session.DropletIP)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateSessionResponse{
		SessionID:       session.ID,
		WireGuardConfig: clientConfig,
	})
}
