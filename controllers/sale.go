package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func SaleIndex(c fiber.Ctx) error {
	var sales []*models.Sale

	if res := config.DB.Debug().Find(&sales); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "sale not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": sales,
	})
}

func SaleShow(c fiber.Ctx) error {
	var sale []*models.Sale

	if result := config.DB.Debug().First(&sale, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "sale not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": sale,
	})

}

func SaleCreate(c fiber.Ctx) error {
	sale := new(models.Sale)

	if err := c.Bind().Body(sale); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if sale.SaleNo == "" ||
		sale.SaleDate.GoString() == "" ||
		sale.Customer == "" ||
		sale.ShippingCost < 0 ||
		sale.Tax1 < 0 ||
		sale.Tax2 < 0 ||
		sale.Total < 0 ||
		sale.AccountID == "" ||
		sale.PaymentType == "" ||
		sale.Reference == "" ||
		sale.Notes == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(sale)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newsale := models.Sale{

		ID:            uuid.New(),
		SaleNo:        sale.SaleNo,
		SaleDate:      sale.SaleDate,
		Customer:      sale.Customer,
		ShippingCost:  sale.ShippingCost,
		Tax1:          sale.Tax1,
		Tax2:          sale.Tax2,
		Total:         sale.Total,
		AccountID:     sale.AccountID,
		PaymentType:   sale.PaymentType,
		Reference:     sale.Reference,
		Notes:         sale.Notes,
		PaymentStatus: sale.PaymentStatus,
	}

	config.DB.Debug().Create(&newsale)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": sale,
	})
}

func SaleUpdate(c fiber.Ctx) error {
	sale := new(models.Sale)

	if err := c.Bind().Body(sale); err != nil {
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
	if sale.SaleNo == "" ||
		sale.SaleDate.GoString() == "" ||
		sale.Customer == "" ||
		sale.ShippingCost < 0 ||
		sale.Tax1 < 0 ||
		sale.Tax2 < 0 ||
		sale.Total < 0 ||
		sale.AccountID == "" ||
		sale.PaymentType == "" ||
		sale.Reference == "" ||
		sale.Notes == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Sale{}).Where("id = ?", id).Updates(map[string]interface{}{
		"sale_no":       sale.SaleNo,
		"sale_date":     sale.SaleDate,
		"customer":      sale.Customer,
		"shippingprice": sale.ShippingCost,
		"tax1":          sale.Tax1,
		"tax2":          sale.Tax2,
		"total":         sale.Total,
		"accid":         sale.AccountID,
		"paymenttype":   sale.PaymentType,
		"reference":     sale.Reference,
		"notes":         sale.Notes,
		"status ":       sale.PaymentStatus})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": sale,
	})
}

func SaleDelete(c fiber.Ctx) error {
	sale := new(models.Sale)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&sale)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "sale deleted",
	})
}
