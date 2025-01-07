package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func SupplierRoutes(r fiber.Router) {
	r.Get("/suppliers/:id", controllers.SupplierShow)
	r.Get("/suppliers", controllers.SupplierIndex)
	r.Post("/suppliers", controllers.SupplierCreate)
	r.Put("/suppliers/:id", controllers.SupplierUpdate)
	r.Delete("/suppliers/:id", controllers.SupplierDelete)
}
