package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func CategoryIndex(c fiber.Ctx) error {
	var category []*models.Categories

	if res := config.DB.Debug().Find(&category); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "category not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": category,
	})
}

func CategoryShow(c fiber.Ctx) error {
	var category []*models.Categories

	if result := config.DB.Debug().First(&category, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "category not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": category,
	})

}

func CategoryCreate(c fiber.Ctx) error {
	category := new(models.Categories)

	if err := c.Bind().Body(category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if category.CategoryCode == "" ||
		category.CategoryName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if category exists
	exists, err := helper.CheckCategoryExists(config.DB, category.CategoryCode)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   err,
		})
	}

	if exists {
		return c.Status(400).JSON(fiber.Map{
			"message": "branch Code already registered",
		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(category)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newCategory := models.Categories{

		ID:           uuid.New(),
		CategoryCode: category.CategoryCode,
		CategoryName: category.CategoryName,
	}
	config.DB.Debug().Create(&newCategory)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": category,
	})
}

func CategoryUpdate(c fiber.Ctx) error {
	category := new(models.Categories)

	if err := c.Bind().Body(category); err != nil {
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
	if category.CategoryCode == "" ||
		category.CategoryName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Categories{}).Where("id = ?", id).Updates(map[string]interface{}{
		"category_code": category.CategoryCode,
		"category_name": category.CategoryName,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": category,
	})
}

func CategoryDelete(c fiber.Ctx) error {
	category := new(models.Categories)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&category)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted",
	})
}

/* func categoryearch(c fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */
