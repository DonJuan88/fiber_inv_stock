package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v3"
)

func BranchStockRoutes(r fiber.Router) {
	//	r.Get("/branchStock/:id", controllers.BranchStockShow)
	r.Get("/branchStock", controllers.BranchStockIndex)
	r.Post("/branchStock", controllers.BranchStockCreate)
	r.Put("/branchStock/:id", controllers.BranchStockUpdate)
	//	r.Delete("/branchStock/:id", controllers.BranchStockDelete)
}
