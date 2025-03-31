package main

import (
	"github.com/rs/cors"
	"license/auth"
	"license/config"
	"license/handlers"
	"license/middleware"
	"net/http"
)

func main() {
	// Connect to db
	config.Connect()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // Adjust to your frontend URL
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/check-token", auth.CheckTokenHandler)
	http.HandleFunc("/create", middleware.AuthenticateMiddleware(handlers.AddKey))
	http.HandleFunc("/delete", middleware.AuthenticateMiddleware(handlers.DeleteKey))
	http.HandleFunc("/authenticate", middleware.AuthenticateMiddleware(handlers.AuthenticateKey))

	// Listen
	http.ListenAndServe(":8090", corsHandler.Handler(http.DefaultServeMux))

}
