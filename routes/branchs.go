package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func BranchRoutes(r fiber.Router) {
	r.Get("/branchs/:id", controllers.BranchShow)
	r.Get("/branchs", controllers.BranchIndex)
	r.Post("/branchs", controllers.BranchCreate)
	r.Put("/branchs/:id", controllers.BranchUpdate)
	r.Delete("/branchs/:id", controllers.BranchDelete)
}
