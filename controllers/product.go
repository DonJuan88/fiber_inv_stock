package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ProductIndex(c fiber.Ctx) error {
	var products []*models.Product

	if res := config.DB.Debug().Find(&products); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "product not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": products,
	})
}

func ProductShow(c fiber.Ctx) error {
	var product []*models.Product

	if result := config.DB.Debug().First(&product, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": product,
	})

}

func ProductCreate(c fiber.Ctx) error {
	product := new(models.Product)

	if err := c.Bind().Body(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if product.ProductCode == "" ||
		product.Barcode1 == "" ||
		product.Barcode2 == "" ||
		product.ProductName == "" ||
		product.Description == "" ||
		product.Category == "" ||
		product.Brand == "" ||
		product.BasePrice < 0 ||
		product.SalePrice < 0 ||
		product.Unit == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if product exists
	exists, err := helper.CheckProductExists(config.DB, product.ProductCode)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   err,
		})
	}

	if exists {
		return c.Status(400).JSON(fiber.Map{
			"message": "Product Code already registered",
		})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(product)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newproduct := models.Product{

		ID:          uuid.New(),
		ProductCode: product.ProductCode,
		Barcode1:    product.Barcode1,
		Barcode2:    product.Barcode2,
		ProductName: product.ProductName,
		Description: product.Description,
		Category:    product.Category,
		Brand:       product.Brand,
		BasePrice:   product.BasePrice,
		SalePrice:   product.SalePrice,
		Unit:        product.Unit,
		Active:      product.Active}
	//
	// hashPassword, err := utils.HashPassword(user.Password)
	// if err != nil {
	//  return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//   "message": "status internal server error",
	//  })
	// }
	//
	//newUser.Password = hashPassword

	config.DB.Debug().Create(&newproduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Created new product",
	})
}

func ProductUpdate(c fiber.Ctx) error {
	product := new(models.Product)

	if err := c.Bind().Body(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// id, _ := strconv.Atoi(c.Params("id"))

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	// Validate required fields
	if product.ProductCode == "" ||
		product.Barcode1 == "" ||
		product.Barcode2 == "" ||
		product.ProductName == "" ||
		product.Description == "" ||
		product.Category == "" ||
		product.Brand == "" ||
		product.BasePrice < 0 ||
		product.SalePrice < 0 ||
		product.Unit == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Product{}).Where("id = ?", id).Updates(map[string]interface{}{
		"code":      product.ProductCode,
		"barcode1":  product.Barcode1,
		"barcode2":  product.Barcode2,
		"name":      product.ProductName,
		"desc":      product.Description,
		"category":  product.Category,
		"brand":     product.Brand,
		"baseprice": product.BasePrice,
		"saleprice": product.SalePrice,
		"unit":      product.Unit,
		"active":    product.Active,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": product,
	})
}

func ProductDelete(c fiber.Ctx) error {
	products := new(models.Product)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&products)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted",
	})
}
