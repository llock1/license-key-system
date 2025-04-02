package main

import (
	"fmt"
	"license/config"
	"license/database"
	"license/middleware"
	"license/routes"

	"github.com/alexflint/go-arg"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	ALLOWED_ORIGINS := fmt.Sprintf("%s", config.Vars.FrontendURL)
	app.Use(cors.New(cors.Config{
		AllowOrigins: ALLOWED_ORIGINS,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// UNRESTRICTED VIEWS
	app.Post("/api/auth", routes.AuthUser)
	app.Post("/api/check-token", routes.CheckTokenHandler)

	jwtSecret := []byte(config.Vars.JWTSecret)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: jwtSecret},
	}))
	app.Use(middleware.AuthMiddleware())

	// RESTRICTED VIEWS
	app.Get("/api/keys", routes.AllKeys)
	app.Get("/api/keys/:id/delete", routes.DeleteKey)
	app.Get("/api/restricted", routes.RestrictedExample)

	fmt.Printf("Listening on port %s\n", config.Vars.Port)
	fmt.Printf("Accepting requests from %s\n", config.Vars.FrontendURL)

	app.Listen(fmt.Sprintf(":%s", config.Vars.Port))
}
