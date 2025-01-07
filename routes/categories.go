package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func CategoryRoutes(r fiber.Router) {
	r.Get("/category/:id", controllers.CategoryShow)
	r.Get("/category", controllers.CategoryIndex)
	r.Post("/category", controllers.CategoryCreate)
	r.Put("/category/:id", controllers.CategoryUpdate)
	r.Delete("/category/:id", controllers.CategoryDelete)
}
