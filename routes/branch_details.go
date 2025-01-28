package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func BranchStockRoutes(r fiber.Router) {
	//	r.Get("/branchStock/:id", controllers.BranchStockShow)
	r.Get("/branchstock", controllers.BranchStockIndex)
	r.Post("/branchstock", controllers.BranchStockCreate)
	r.Put("/branchstock/:id", controllers.BranchStockUpdate)
	//	r.Delete("/branchStock/:id", controllers.BranchStockDelete)
}
