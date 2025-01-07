package routes

import (
	"inv_fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func NotificationRoutes(r fiber.Router) {
	r.Get("/notifications/:id", controllers.NotificationShow)
	r.Get("/notifications", controllers.NotificationIndex)
	r.Post("/notifications", controllers.NotificationCreate)
	r.Put("/notifications/:id", controllers.NotificationUpdate)
	r.Delete("/notifications/:id", controllers.NotificationDelete)
}