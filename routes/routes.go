package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sergio-rey/jwt-go/controller"
)

func Setup(app *fiber.App) {

	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
	app.Get("/user", controller.User)
	app.Post("/logout", controller.Logout)

}
