package models

import "time"

type Session struct {
	ID            string
	Region        string
	DropletID     int
	DropletIP     string
	Status        string
	CreatedAt     time.Time
	ServerKeyPair *KeyPair
	ClientKeyPair *KeyPair
}
