package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Adi-111/spotifyDev/internal/routes"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Initializing.....")
	r := mux.NewRouter() // creating router with gorilla-mux
	fmt.Println("Router Created.....")
	routes.RegisterRoutes(r) // registering the routes to server
	staticDir := "../static"
	fileServer := http.FileServer(http.Dir(staticDir))
	r.PathPrefix("/").Handler(fileServer)

	fmt.Println("Server is starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
