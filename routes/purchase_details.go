package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func PurchaseDetailRoutes(r fiber.Router) {
	r.Get("/purchasedetails/:id", controllers.PurchaseDetailShow)
//	r.Get("/purchasedetails", controllers.PurchaseDetailIndex)
	r.Post("/purchasedetails", controllers.PurchaseDetailCreate)
	r.Put("/purchasedetails/:id", controllers.PurchaseDetailUpdate)
	r.Delete("/purchasedetails/:id", controllers.PurchaseDetailDelete)
}