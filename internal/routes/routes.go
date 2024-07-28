package routes

import (
	"fmt"
	"net/http"

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
	r.HandleFunc("/success", handlers.SuccessHandler).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")

	// Serve static files at the root path or a specific path (e.g., /static/)
	staticDir := "static/"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	// Set the root route to serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

}
