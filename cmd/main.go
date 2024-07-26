package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Adi-111/spotifyDev/internal/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Initializing.....")
	r := mux.NewRouter() // creating router with gorilla-mux
	fmt.Println("Router Created.....")
	routes.RegisterRoutes(r) // registering the routes to server
	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:4200"}))(r)

	fmt.Println("Server is starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", corsHandler))

}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
