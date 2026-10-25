package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// healthHandler handles GET /health
// Returns JSON: {"status":"ok","service":"backend-from-first-principles"}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Map that will be converted to JSON
	resp := map[string]string{
		"status":  "ok",
		"service": "backend-from-first-principles",
	}

	// Encode map to JSON and write to response
	json.NewEncoder(w).Encode(resp)
}

// getUserHandler handles GET /users/{id}
// {id} comes from the URL path
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the {id} value from the URL path
	id := r.PathValue("id")

	// Hardcoded user data (no database yet)
	resp := map[string]interface{}{
		"id":    id,
		"name":  "Yatan",
		"email": "yatan@jat.com",
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Create a new ServeMux (Go's request router, like Spring's DispatcherServlet)
	mux := http.NewServeMux()

	// Register routes
	// GET /health -> healthHandler
	mux.HandleFunc("GET /health", healthHandler)

	// GET /users/{id} -> getUserHandler
	mux.HandleFunc("GET /users/{id}", getUserHandler)

	// Start server on port 8080
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
