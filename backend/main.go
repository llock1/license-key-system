package main

import (
	"license/auth"
	"license/config"
	"license/database"
	"license/handlers"
	"license/middleware"
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/rs/cors"
)

func main() {

	var args struct {
		Development bool
	}

	arg.MustParse(&args)

	if !config.InitConfig(args.Development) {
		log.Fatal("failed to initialize config")
	}

	// Connect to db
	database.Connect()

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
