package main

import (
	"fmt"
	"license/config"
	"license/database"
	"license/middleware"
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
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// UNRESTRICTED VIEWS
	app.Post("/api/login", routes.LoginUser)
	app.Post("/api/register", routes.RegisterUser)
	app.Post("/api/check-token", routes.CheckTokenHandler)

	app.Use(middleware.AuthMiddleware())

	// RESTRICTED VIEWS
	app.Get("/api/licenses", routes.GetLicenses)
	app.Delete("/api/licenses/:id", routes.DeleteLicense)
	app.Post("/api/licenses", routes.CreateLicense)

	app.Listen(fmt.Sprintf(":%s", config.Vars.Port))
}
