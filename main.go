package main

import (
	"fabopgg/helpers"
	"net/http"
)

func main() {
	// Initialize the database
	helpers.InitDB()

	// Create a new ServeMux and register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/add-match", helpers.AddMatch)
	mux.HandleFunc("/get-matches", helpers.GetMatches)

	// Wrap the router with the CORS middleware
	corsMux := helpers.CorsMiddleware(mux)

	// Start the server with the CORS-enabled handler
	http.ListenAndServe(":8080", corsMux)
}
