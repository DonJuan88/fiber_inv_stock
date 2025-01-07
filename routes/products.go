package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func ProductRoutes(r fiber.Router) {
	r.Get("/products/:id", controllers.ProductShow)
	r.Get("/products", controllers.ProductIndex)
	r.Post("/products", controllers.ProductCreate)
	r.Put("/products/:id", controllers.ProductUpdate)
	r.Delete("/products/:id", controllers.ProductDelete)
}
