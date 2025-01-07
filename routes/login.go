package routes

import (
	"inv_fiber/middleware"

	"github.com/gofiber/fiber/v3"
)

func LoginRoutes(r fiber.Router) {
	//	r.Get("/login/:id", controllers.loginShow)
	r.Post("/login", middleware.UserLogin)
	//	r.Post("/login", controllers.loginPost)
	////	r.Put("/login/:id", controllers.Ima)
	//	r.Delete("/login/:id", controllers.loginDelete)
}
