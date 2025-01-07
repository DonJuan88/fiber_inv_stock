package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func CompanyRoutes(r fiber.Router) {
	r.Get("/company/:id", controllers.CompanyShow)
//	r.Get("/company", controllers.CompanyIndex)
	r.Post("/company", controllers.CompanyCreate)
	r.Put("/company/:id", controllers.CompanyUpdate)
	r.Delete("/company/:id", controllers.CompanyDelete)
}