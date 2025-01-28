package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func ImageRoutes(r fiber.Router) {
	r.Get("/images/:id", controllers.ImageShow)
	r.Get("/images", controllers.ImageIndex)
	r.Post("/images", controllers.ImagePost)
	//	r.Put("/image/:id", controllers.Ima)
	r.Delete("/images/:id", controllers.ImageDelete)
}
