package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/SamuelMR98/asset-mgmt-dApp/eth"
	"github.com/SamuelMR98/asset-mgmt-dApp/models"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Handler struct {
	PgPool      *pgxpool.Pool
	RedisClient *redis.Client
	EthClient   *eth.EthClient
}

func NewHandler(pgPool *pgxpool.Pool, redisClient *redis.Client, ethClient *eth.EthClient) *Handler {
	return &Handler{
		PgPool:      pgPool,
		RedisClient: redisClient,
		EthClient:   ethClient,
	}
}

// GetUsers returns a list of users (demo using fake data)
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
		{ID: 1, Username: "alice", Role: "admin"},
		{ID: 2, Username: "bob", Role: "user"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser inserts a new user (demo: echoes back provided data)
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// In production, hash the password and save the user to the database.
	user.ID = 999 // dummy ID for demo
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetAssets returns a list of assets (demo with fake data)
func (h *Handler) GetAssets(w http.ResponseWriter, r *http.Request) {
	assets := []models.Asset{
		{ID: 1, Name: "Digital Gold", Type: "Token", OwnerID: 1},
		{ID: 2, Name: "Crypto Art", Type: "NFT", OwnerID: 2},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
}

// CreateAsset creates a new asset and simulates deploying a smart contract.
func (h *Handler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var asset models.Asset
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	asset.ID = 100 // dummy asset ID

	// Simulate smart contract deployment for this asset.
	contractAddress, err := h.EthClient.DeploySmartContract()
	if err != nil {
		http.Error(w, "Error deploying smart contract", http.StatusInternalServerError)
		return
	}
	log.Printf("Deployed smart contract at address: %s", contractAddress)

	response := map[string]interface{}{
		"asset":         asset,
		"smartContract": contractAddress,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetTransactions returns a demo list of transactions.
func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions := []models.Transaction{
		{ID: 1, AssetID: 1, FromUserID: 1, ToUserID: 2, Timestamp: time.Now(), Amount: 10.0},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction creates a new transaction record (demo).
func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var txn models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&txn); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	txn.ID = 200 // dummy transaction ID
	txn.Timestamp = time.Now()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(txn)
}

// GetMarketData fetches simulated market data using Redis caching.
func (h *Handler) GetMarketData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cacheKey := "marketdata"
	// Try to fetch data from Redis first.
	data, err := h.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// Cache miss: simulate fetching market data.
		marketData := map[string]interface{}{
			"BTC": 50000.0,
			"ETH": 4000.0,
			"LTC": 180.0,
		}
		jsonBytes, err := json.Marshal(marketData)
		if err != nil {
			http.Error(w, "Error marshalling data", http.StatusInternalServerError)
			return
		}
		// Cache the data for 60 seconds.
		h.RedisClient.Set(ctx, cacheKey, jsonBytes, 60*time.Second)
		data = string(jsonBytes)
	} else if err != nil {
		http.Error(w, "Error fetching market data from cache", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}
