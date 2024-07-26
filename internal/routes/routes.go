package routes

import (
	"fmt"

	"github.com/Adi-111/spotifyDev/internal/config"
	"github.com/Adi-111/spotifyDev/internal/handlers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(r *mux.Router) {
	fmt.Println("routing...")
	config.ConnectDB()
	fmt.Println("DB Online...")

	config.MigrateDB()
	fmt.Println("DB Ready to use...")
	r.HandleFunc("/login", handlers.SpotifyLoginHandler).Methods("GET")
	r.HandleFunc("/callback", handlers.CallbackHandler).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")

}
