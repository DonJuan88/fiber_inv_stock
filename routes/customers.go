package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func CustomerRoutes(r fiber.Router) {
	r.Get("/customer/:id", controllers.CustomerShow)
	r.Get("/customer", controllers.CustomerIndex)
	r.Post("/customer", controllers.CustomerCreate)
	r.Put("/customer/:id", controllers.CustomerUpdate)
	r.Delete("/customer/:id", controllers.CustomerDelete)
}
