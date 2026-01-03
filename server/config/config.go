package config

import (
    "os"
    "github.com/digitalocean/godo"
    "github.com/joho/godotenv"
)

func LoadDOClient() *godo.Client {
    godotenv.Load()
    return godo.NewFromToken(os.Getenv("DO_TOKEN"))
}
