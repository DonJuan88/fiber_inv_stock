package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router) {
	r.Get("/users/:id", controllers.UserShow)
	r.Get("/users", controllers.UserIndex)
	r.Post("/users", controllers.UserCreate)
	r.Put("/users/:id", controllers.UserUpdate)
	r.Delete("/users/:id", controllers.UserDelete)
}