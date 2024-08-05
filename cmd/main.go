package main

import (
	"log"
	"net/http"

	"github.com/balasl342/api-security/internal/handlers"
	"github.com/balasl342/api-security/internal/middleware"
	"github.com/balasl342/api-security/pkg/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Redis
	redis.Init()

	// Create a new router
	r := mux.NewRouter()

	// Public endpoints
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Apply middleware to protected routes
	protected := r.NewRoute().Subrouter()
	protected.Use(middleware.Authenticate)
	protected.Use(middleware.RateLimiter)
	protected.HandleFunc("/current-time", handlers.CurrentTimeHandler).Methods("GET")
	protected.HandleFunc("/echo", handlers.EchoHandler).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}
