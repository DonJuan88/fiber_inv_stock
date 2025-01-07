package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func SupplierIndex(c *fiber.Ctx) error {
	var suppliers []*models.Supplier

	if res := config.DB.Debug().Find(&suppliers); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "supplier not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": suppliers,
	})
}
func SupplierShow(c *fiber.Ctx) error {
	var supplier []*models.Supplier

	if result := config.DB.Debug().First(&supplier, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "suppliers not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": supplier,
	})

}

func SupplierCreate(c *fiber.Ctx) error {
	supplier := new(models.Supplier)

	if err := c.BodyParser(supplier); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 
	supplier.Code   == "" ||
	supplier.Name          == "" ||
	supplier.ContactPerson == "" ||
	supplier.Email         == "" ||
	supplier.Phone         == "" ||
	supplier.Address       == "" ||
	supplier.City          == "" ||
	supplier.State         == "" ||
	supplier.PostalCode    == "" ||
	supplier.Country       == "" {
		
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}


	// Check if supplier exists
	exists, err := helper.CheckSupplierExists(config.DB, supplier.Code)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
   "message": "failed to validate",
   "error":  err,
  })
	}

	
	if exists {
		return c.Status(400).JSON(fiber.Map{
	   "message": "branch Code already registered",

		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(supplier)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newsupplier := models.Supplier{

		ID: uuid.New(),
		Code :supplier.Code,
		Name:supplier.Name,
		ContactPerson:supplier.ContactPerson,
		Email:supplier.Email,
		Phone: supplier.Phone,
		Address:supplier.Address,
		City:supplier.City,
		State:supplier.State,
		PostalCode:supplier.PostalCode,
		Country: supplier.Country,
	}
	config.DB.Debug().Create(&newsupplier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": supplier,
	})
}

func SupplierUpdate(c *fiber.Ctx) error {
	supplier := new(models.Supplier)

	if err := c.BodyParser(supplier); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	// Validate required fields
	if 
	supplier.Code   == "" ||
	supplier.Name          == "" ||
	supplier.ContactPerson == "" ||
	supplier.Email         == "" ||
	supplier.Phone         == "" ||
	supplier.Address       == "" ||
	supplier.City          == "" ||
	supplier.State         == "" ||
	supplier.Country       == "" {
		
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Supplier{}).Where("id = ?", id).Updates(map[string]interface{}{
		"code":supplier.Code,
		"name":supplier.Name,
		"cp":supplier.ContactPerson,
		"email":supplier.Email,
		"phone":supplier.Phone,
		"address":supplier.Address,
		"city":supplier.City,
		"state":supplier.State,
		"postalcode":supplier.PostalCode,
		"country":supplier.Country,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": supplier,
	})
}

func SupplierDelete(c *fiber.Ctx) error {
	supplier := new(models.Supplier)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&supplier)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "supplier deleted",
	})
}

/* func Supplierearch(c *fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */