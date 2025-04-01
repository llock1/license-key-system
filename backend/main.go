package main

import (
	"fmt"
	"license/config"
	"license/database"
	"license/routes"

	"github.com/alexflint/go-arg"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	//"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	var args struct {
		Development bool
	}

	arg.MustParse(&args)

	config.Initialize(args.Development)

	database.Connect()

	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	//app.Use(cors.New(cors.Config{
	//	AllowOrigins: []string{config.Vars.FrontendURL},
	//	AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
	//}))

	app.Post("/api/auth", routes.AuthUser)
	app.Get("/api/restricted", routes.RestrictedExample)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	fmt.Printf("Listening on port %s\n", config.Vars.Port)

	app.Listen(fmt.Sprintf(":%s", config.Vars.Port))
}
