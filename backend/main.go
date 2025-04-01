package main

import (
	"fmt"
	"license/auth"
	"license/config"
	"license/database"
	"license/handlers"
	"license/middleware"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/rs/cors"
)

func main() {

	var args struct {
		Development bool
	}

	arg.MustParse(&args)

	config.Initialize(args.Development)

	database.Connect()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{config.Vars.FrontendURL},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/check-token", auth.CheckTokenHandler)
	http.HandleFunc("/create", middleware.AuthenticateMiddleware(handlers.AddKey))
	http.HandleFunc("/delete", middleware.AuthenticateMiddleware(handlers.DeleteKey))
	http.HandleFunc("/authenticate", middleware.AuthenticateMiddleware(handlers.AuthenticateKey))

	fmt.Printf("Listening on port %s\n", config.Vars.Port)

	http.ListenAndServe(fmt.Sprintf(":%s", config.Vars.Port), corsHandler.Handler(http.DefaultServeMux))
}
