package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func BrandIndex(c *fiber.Ctx) error {
	var brands []*models.Brands

	if res := config.DB.Debug().Find(&brands); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Brand not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": brands,
	})
}

func BrandShow(c *fiber.Ctx) error {
	var brand []*models.Brands

	if result := config.DB.Debug().First(&brand, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Brand not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": brand,
	})

}

func BrandCreate(c *fiber.Ctx) error {
	brand := new(models.Brands)

	if err := c.BodyParser(brand); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 	brand.BrandCode == "" ||
		brand.BrandName == ""  {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if Brand exists
	exists, err := helper.CheckBrandExists(config.DB, brand.BrandCode)
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
	errValidate := validate.Struct(brand)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newBrand := models.Brands{

		ID:          uuid.New(),
		BrandCode:  brand.BrandCode,
		BrandName: brand.BrandName,
	}
	config.DB.Debug().Create(&newBrand)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": brand,
		
	})
}

func BrandUpdate(c *fiber.Ctx) error {
	brand := new(models.Brands)

	if err := c.BodyParser(brand); err != nil {
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
	if 	brand.BrandCode == "" ||
		brand.BrandName == ""  {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}
	

	config.DB.Debug().Model(&models.Brands{}).Where("id = ?", id).Updates(map[string]interface{}{
		"brand_code" : brand.BrandCode,
		"brand_name" : brand.BrandName,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": brand,
	})
}

func BrandDelete(c *fiber.Ctx) error {
	brands := new(models.Brands)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&brands)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Brand deleted",
	})
}

/* func BrandSearch(c *fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */