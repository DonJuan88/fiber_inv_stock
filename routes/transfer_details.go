package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func TransferDetailRoutes(r fiber.Router) {
	r.Get("/transferdetails/:id", controllers.TransferDetailShow)
	//r.Get("/transferdetails", controllers.TransferDetailIndex)
	r.Post("/transferdetails", controllers.TransferDetailCreate)
	r.Put("/transferdetails/:id", controllers.TransferDetailUpdate)
	r.Delete("/transferdetails/:id", controllers.TransferDetailDelete)
}