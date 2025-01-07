package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func SaleRoutes(r fiber.Router) {
	r.Get("/sales/:id", controllers.SaleShow)
	r.Get("/sales", controllers.SaleIndex)
	r.Post("/sales", controllers.SaleCreate)
	r.Put("/sales/:id", controllers.SaleUpdate)
	r.Delete("/sales/:id", controllers.SaleDelete)
}
