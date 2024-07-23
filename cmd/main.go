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
	r := mux.NewRouter()
	fmt.Println("Router Created.....")

	routes.RegisterRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
	fmt.Println("online on http://localhost:8080")

}
