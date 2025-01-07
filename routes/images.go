package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(r fiber.Router) {
	r.Get("/image/:id", controllers.ImageShow)
	r.Get("/image", controllers.ImageIndex)
	r.Post("/image", controllers.ImagePost)
//	r.Put("/image/:id", controllers.Ima)
	r.Delete("/image/:id", controllers.ImageDelete)
}