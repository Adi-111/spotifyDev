package main

import (
	"fmt"
	"net/http"

	"github.com/Adi-111/spotifyDev/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Initializing.....")
	r := mux.NewRouter() // creating router with gorilla-mux
	fmt.Println("Router Created.....")
	routes.RegisterRoutes(r)
	// registering the routes to server

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
