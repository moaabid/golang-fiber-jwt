package main

import (
	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/moaabid/golang-fiber-jwt/data"
	"github.com/moaabid/golang-fiber-jwt/routes"
)

func main() {

	app := fiber.New()
	var err error
	data.Engine, err = data.CreateDBEngine()

	if err != nil {
		panic(err)
	}

	setupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}

func setupRoutes(app *fiber.App) {
	app.Post("/signup", routes.SignUp)
	app.Post("/login", routes.Login)

	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Get("/", routes.Private)

	public := app.Group("/public")
	public.Get("/", routes.Public)
}
