package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func PurchaseRoutes(r fiber.Router) {
	r.Get("/purchases/:id", controllers.PurchaseShow)
	r.Get("/purchases", controllers.PurchaseIndex)
	r.Post("/purchases", controllers.PurchaseCreate)
	r.Put("/purchases/:id", controllers.PurchaseUpdate)
	r.Delete("/purchases/:id", controllers.PurchaseDelete)
}
