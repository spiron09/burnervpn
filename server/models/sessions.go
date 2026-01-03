package models

type Session struct {
	ID            string
	Region        string
	DropletID     string
	DropletIP     string
	Status        string
	CreatedAt     string
	ServerKeyPair *KeyPair
	ClientKeyPair *KeyPair
}
