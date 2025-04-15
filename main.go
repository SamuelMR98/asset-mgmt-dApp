package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SamuelMR98/asset-mgmt-dApp/config"
	"github.com/SamuelMR98/asset-mgmt-dApp/db"
	"github.com/SamuelMR98/asset-mgmt-dApp/eth"
	"github.com/SamuelMR98/asset-mgmt-dApp/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize PostgreSQL connection
	pgPool, err := db.InitPostgres(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to Postgres: %v", err)
	}
	defer pgPool.Close()

	// Initialize Redis client
	redisClient, err := db.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize Ethereum client (for blockchain operations)
	ethClient, err := eth.InitEthereumClient(cfg.EthereumNodeURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}
	// Create HTTP router
	router := mux.NewRouter()

	// Create handler instance with dependencies
	h := handlers.NewHandler(pgPool, redisClient, ethClient)

	// Define API routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/users", h.GetUsers).Methods("GET")
	api.HandleFunc("/users", h.CreateUser).Methods("POST")
	api.HandleFunc("/assets", h.GetAssets).Methods("GET")
	api.HandleFunc("/assets", h.CreateAsset).Methods("POST")
	api.HandleFunc("/transactions", h.GetTransactions).Methods("GET")
	api.HandleFunc("/transactions", h.CreateTransaction).Methods("POST")
	api.HandleFunc("/marketdata", h.GetMarketData).Methods("GET")

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
