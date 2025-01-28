package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func CustomerRoutes(r fiber.Router) {
	r.Get("/customers/:id", controllers.CustomerShow)
	r.Get("/customers", controllers.CustomerIndex)
	r.Post("/customers", controllers.CustomerCreate)
	r.Put("/customers/:id", controllers.CustomerUpdate)
	r.Delete("/customers/:id", controllers.CustomerDelete)
}
