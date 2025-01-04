package main

import (
	"inv_fiber/config"
	"inv_fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()
	config.DatabaseConnection()

	app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Inventory API with fiber",
})



	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api/v1")

	// Register rute modular
	//routes.AccountUserRoutes(api) // Anda bisa menambahkan fungsi ini di file routes/account_user_routes.go
	routes.BranchRoutes(api)
//	routes.BrandRoutes(api) // And

	app.Listen(":3000")
}