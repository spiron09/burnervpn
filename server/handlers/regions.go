package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/digitalocean/godo"
	"github.com/spiron09/burnervpn/server/config"
	"github.com/spiron09/burnervpn/server/models"
)

func HandleRegions(w http.ResponseWriter, r *http.Request) {
	client := config.LoadDOClient()

	regions, _, err := client.Regions.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	availableRegions := []models.Region{}
	for _, region := range regions {
		if region.Available {
			availableRegions = append(availableRegions, models.Region{
				Slug: region.Slug,
				Name: region.Name,
			})
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(availableRegions)
}
