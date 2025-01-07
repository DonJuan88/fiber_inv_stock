package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func MessageRoutes(r fiber.Router) {
	r.Get("/message/:id", controllers.MessageShow)
	//	r.Get("/message", controllers.MessageIndex)
	r.Post("/message", controllers.MessageCreate)
	//	r.Put("/message/:id", controllers.Ima)
	r.Delete("/message/:id", controllers.MessageDelete)
}
