package main

import (
	"inv_fiber/config"
	"inv_fiber/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	config.LoadConfig()
	config.DatabaseConnection()

	app := fiber.New(
	//fiber.Config{
	//CaseSensitive: true,
	//StrictRouting: true,
	//ServerHeader:  "Fiber",
	//AppName:       "Inventory API with fiber",}
	)

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     []string{"*"},
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		},
		AllowHeaders:        []string{},
		AllowCredentials:    false,
		ExposeHeaders:       []string{},
		MaxAge:              0,
		AllowPrivateNetwork: false,
	}))

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	routes.LoginRoutes(api)
	v1 := api.Group("/v1")
	// Register rute modular
	//routes.AccountUserRoutes(v1) // Anda bisa menambahkan fungsi ini di file routes/account_user_routes.go
	routes.BranchRoutes(v1)
	routes.BranchRoutes(v1)
	routes.BrandRoutes(v1)
	routes.CategoryRoutes(v1)
	routes.CompanyRoutes(v1)
	routes.CustomerRoutes(v1)
	routes.ImageRoutes(v1)
	routes.MessageRoutes(v1)
	routes.NotificationRoutes(v1)
	routes.SaleRoutes(v1)
	routes.SaleDetailRoutes(v1)
	routes.ProductRoutes(v1)
	routes.PurchaseDetailRoutes(v1)
	routes.PurchaseRoutes(v1)
	routes.SupplierRoutes(v1)
	routes.TransferDetailRoutes(v1)
	routes.TransferRoutes(v1)
	routes.UserRoutes(api)
	//	routes.BrandRoutes(v1) // And

	app.Listen(":3000")
}
