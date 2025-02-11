package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

/*
	 func purchaseDetailIndex(c fiber.Ctx) error {
		var purchaseDetails []*models.purchaseDetailetail

		if res := config.DB.Debug().Find(&purchaseDetails); res.Error != nil {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "purchaseDetail not found",
			})

		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": purchaseDetails,
		})
	}
*/
func PurchaseDetailShow(c fiber.Ctx) error {
	var purchaseDetail []*models.PurchaseDetail

	if result := config.DB.Debug().First(&purchaseDetail, c.Params("id=?")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "purchase not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchaseDetail,
	})

}

func PurchaseDetailCreate(c fiber.Ctx) error {
	purchaseDetail := new(models.PurchaseDetail)

	if err := c.Bind().Body(purchaseDetail); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if purchaseDetail.PurchaseNo == "" ||
		purchaseDetail.PurchaseDate.GoString() == "" ||
		purchaseDetail.ProductCode == "" ||
		purchaseDetail.Qty < 0 ||
		purchaseDetail.Price < 0 ||
		purchaseDetail.Discount < 0 ||
		purchaseDetail.PurchasePrice < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(purchaseDetail)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newpurchaseDetail := models.PurchaseDetail{

		ID:            uuid.New(),
		PurchaseNo:    purchaseDetail.PurchaseNo,
		PurchaseDate:  purchaseDetail.PurchaseDate,
		ProductCode:   purchaseDetail.ProductCode,
		Qty:           purchaseDetail.Qty,
		Price:         purchaseDetail.Price,
		Discount:      purchaseDetail.Discount,
		PurchasePrice: purchaseDetail.PurchasePrice}
	config.DB.Debug().Create(&newpurchaseDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchaseDetail,
	})
}

func PurchaseDetailUpdate(c fiber.Ctx) error {
	purchaseDetail := new(models.PurchaseDetail)

	if err := c.Bind().Body(purchaseDetail); err != nil {
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
	if purchaseDetail.PurchaseNo == "" ||
		purchaseDetail.PurchaseDate.GoString() == "" ||
		purchaseDetail.ProductCode == "" ||
		purchaseDetail.Qty < 0 ||
		purchaseDetail.Price < 0 ||
		purchaseDetail.Discount < 0 ||
		purchaseDetail.PurchasePrice < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.PurchaseDetail{}).Where("id = ?", id).Updates(map[string]interface{}{
		"purchase_no":   purchaseDetail.PurchaseNo,
		"purchase_date": purchaseDetail.PurchaseDate,
		"product_code":  purchaseDetail.ProductCode,
		"qty":           purchaseDetail.Qty,
		"baseprice":     purchaseDetail.Price,
		"discount":      purchaseDetail.Discount,
		"purchaseprice": purchaseDetail.PurchasePrice})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchaseDetail,
	})
}

func PurchaseDetailDelete(c fiber.Ctx) error {
	purchaseDetail := new(models.PurchaseDetail)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&purchaseDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "purchaseDetail deleted",
	})
}
