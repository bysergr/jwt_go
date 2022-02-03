package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sergio-rey/jwt/database"
	"github.com/sergio-rey/jwt/routes"
)

func main() {

	app := fiber.New()

	database.Connection()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
