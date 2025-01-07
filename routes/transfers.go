package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func TransferRoutes(r fiber.Router) {
	r.Get("/transfers/:id", controllers.TransferShow)
	r.Get("/transfers", controllers.TransferIndex)
	r.Post("/transfers", controllers.TransferCreate)
	r.Put("/transfers/:id", controllers.TransferUpdate)
	r.Delete("/transfers/:id", controllers.TransferDelete)
}
