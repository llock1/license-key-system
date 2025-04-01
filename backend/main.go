package main

import (
	"fmt"
	"license/config"
	"license/database"
	"license/routes"

	"github.com/alexflint/go-arg"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.Vars.FrontendURL},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization"},
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	app.Post("/api/auth", routes.AuthUser)
	app.Post("/api/logout", routes.LogoutUser)

	fmt.Printf("Listening on port %s\n", config.Vars.Port)

	app.Listen(fmt.Sprintf(":%s", config.Vars.Port))
}
