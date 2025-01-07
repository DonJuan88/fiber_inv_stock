package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func BrandRoutes(r fiber.Router) {
	r.Get("/brands/:id", controllers.BrandShow)
	r.Get("/brands", controllers.BrandIndex)
	r.Post("/brands", controllers.BrandCreate)
	r.Put("/brands/:id", controllers.BrandUpdate)
	r.Delete("/brands/:id", controllers.BrandDelete)
}