package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func PurchaseIndex(c fiber.Ctx) error {
	var purchases []*models.Purchase

	if res := config.DB.Debug().Find(&purchases); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "purchase not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchases,
	})
}

func PurchaseShow(c fiber.Ctx) error {
	var purchase []*models.Purchase

	if result := config.DB.Debug().First(&purchase, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "purchase not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchase,
	})

}

func PurchaseCreate(c fiber.Ctx) error {
	purchase := new(models.Purchase)

	if err := c.Bind().Body(purchase); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if purchase.PurchaseNo == "" ||
		purchase.PurchaseDate.GoString() == "" ||
		purchase.BranchCode == "" ||
		purchase.Supplier == "" ||
		purchase.ShippingCost < 0 ||
		purchase.Tax1 < 0 ||
		purchase.Tax2 < 0 ||
		purchase.Total < 0 ||
		purchase.UserID == "" ||
		purchase.PaymentType == "" ||
		purchase.ShipStatus == "" ||
		purchase.Reference == "" ||
		purchase.Notes == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(purchase)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newpurchase := models.Purchase{

		ID:            uuid.New(),
		PurchaseNo:    purchase.PurchaseNo,
		PurchaseDate:  purchase.PurchaseDate,
		BranchCode:    purchase.BranchCode,
		Supplier:      purchase.Supplier,
		ShippingCost:  purchase.ShippingCost,
		Tax1:          purchase.Tax1,
		Tax2:          purchase.Tax2,
		Total:         purchase.Total,
		UserID:        purchase.UserID,
		PaymentType:   purchase.PaymentType,
		ShipStatus:    purchase.ShipStatus,
		Reference:     purchase.Reference,
		Notes:         purchase.Notes,
		PaymentStatus: purchase.PaymentStatus}

	config.DB.Debug().Create(&newpurchase)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchase,
	})
}

func PurchaseUpdate(c fiber.Ctx) error {
	purchase := new(models.Purchase)

	//referenceTimeString := referenceTime.Format(layout)

	if err := c.Bind().Body(purchase); err != nil {
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
	if purchase.PurchaseNo == "" ||
		///		time.Parse(layout, purchase.PurchaseDate) == "" ||
		purchase.BranchCode == "" ||
		purchase.Supplier == "" ||
		purchase.ShippingCost < 0 ||
		purchase.Tax1 < 0 ||
		purchase.Tax2 < 0 ||
		purchase.Total < 0 ||
		purchase.UserID == "" ||
		purchase.PaymentType == "" ||
		purchase.ShipStatus == "" ||
		purchase.Reference == "" ||
		purchase.Notes == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Purchase{}).Where("id = ?", id).Updates(map[string]interface{}{
		"purchase_no":               purchase.PurchaseNo,
		"purchase_date":             purchase.PurchaseDate,
		"branch":                    purchase.BranchCode,
		"supplier":                  purchase.Supplier,
		"shippingprice":             purchase.ShippingCost,
		"tax1":                      purchase.Tax1,
		"tax2":                      purchase.Tax2,
		"total":                     purchase.Total,
		"user_id":                   purchase.UserID,
		"paymenttype":               purchase.PaymentType,
		"shippingstatus":            purchase.ShipStatus,
		"reference":                 purchase.Reference,
		"notes":                     purchase.Notes,
		"status gorm:default:false": purchase.PaymentStatus,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": purchase,
	})
}

func PurchaseDelete(c fiber.Ctx) error {
	purchase := new(models.Purchase)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&purchase)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "purchase deleted",
	})
}
