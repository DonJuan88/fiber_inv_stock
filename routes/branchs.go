package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func BranchRoutes(r fiber.Router) {
	r.Get("/branch/:id", controllers.BranchShow)
	r.Get("/branch", controllers.BranchIndex)
	r.Post("/branch", controllers.BranchCreate)
	r.Put("/branch/:id", controllers.BranchUpdate)
	r.Delete("/branch/:id", controllers.BranchDelete)
}
