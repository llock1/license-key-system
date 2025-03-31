package main

import (
	"license/auth"
	"license/config"
	"license/handlers"
	"license/middleware"
	"net/http"
)

func main() {
	// Connect to db
	config.Connect()

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/create", middleware.AuthenticateMiddleware(handlers.AddKey))
	http.HandleFunc("/delete", middleware.AuthenticateMiddleware(handlers.DeleteKey))
	http.HandleFunc("/authenticate", middleware.AuthenticateMiddleware(handlers.AuthenticateKey))

	// Listen
	http.ListenAndServe(":8090", nil)
}
