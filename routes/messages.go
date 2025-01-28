package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func MessageRoutes(r fiber.Router) {
	r.Get("/messages/:id", controllers.MessageShow)
	//	r.Get("/message", controllers.MessageIndex)
	r.Post("/messages", controllers.MessageCreate)
	//	r.Put("/message/:id", controllers.Ima)
	r.Delete("/messages/:id", controllers.MessageDelete)
}
