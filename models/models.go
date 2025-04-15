package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // password is stored hashed
	Role     string `json:"role"`
}

type Asset struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	OwnerID int    `json:"owner_id"`
}

type Transaction struct {
	ID         int       `json:"id"`
	AssetID    int       `json:"asset_id"`
	FromUserID int       `json:"from_user_id"`
	ToUserID   int       `json:"to_user_id"`
	Timestamp  time.Time `json:"timestamp"`
	Amount     float64   `json:"amount"`
}

type SmartContract struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	AssetID int    `json:"asset_id"`
}
