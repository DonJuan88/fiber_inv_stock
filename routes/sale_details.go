package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func SaleDetailRoutes(r fiber.Router) {
	r.Get("/saledetails/:id", controllers.SaleDetailShow)
	//	r.Get("/saledetails", controllers.SaleDetailIndex)
	r.Post("/saledetails", controllers.SaleDetailCreate)
	r.Put("/saledetails/:id", controllers.SaleDetailUpdate)
	r.Delete("/saledetails/:id", controllers.SaleDetailDelete)
}
