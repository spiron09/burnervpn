package config

import (
	"github.com/digitalocean/godo"
	"github.com/joho/godotenv"
	"os"
)

func LoadDOClient() *godo.Client {
	godotenv.Load()
	return godo.NewFromToken(os.Getenv("DO_TOKEN"))
}
