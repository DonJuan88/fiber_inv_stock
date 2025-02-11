package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

/*
	 func TransferDetailIndex(c fiber.Ctx) error {
		var transferDetails []*models.TransferDetailetail

		if res := config.DB.Debug().Find(&transferDetails); res.Error != nil {
			c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "transferDetail not found",
			})

		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": transferDetails,
		})
	}
*/
func TransferDetailShow(c fiber.Ctx) error {
	var transferDetail []*models.TransferDetail

	if result := config.DB.Debug().Find(&transferDetail, c.Params("id=?")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "transfer not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transferDetail,
	})

}

func TransferDetailCreate(c fiber.Ctx) error {
	transferDetail := new(models.TransferDetail)

	if err := c.Bind().Body(transferDetail); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if transferDetail.TransferNo == "" ||
		transferDetail.ProductCode == "" ||
		transferDetail.Qty < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(transferDetail)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newtransferDetail := models.TransferDetail{

		ID:          uuid.New(),
		TransferNo:  transferDetail.TransferNo,
		ProductCode: transferDetail.ProductCode,
		Qty:         transferDetail.Qty,
	}
	config.DB.Debug().Create(&newtransferDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transferDetail,
	})
}

func TransferDetailUpdate(c fiber.Ctx) error {
	transferDetail := new(models.TransferDetail)

	if err := c.Bind().Body(transferDetail); err != nil {
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
	if transferDetail.TransferNo == "" ||
		transferDetail.ProductCode == "" ||
		transferDetail.Qty < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.TransferDetail{}).Where("id = ?", id).Updates(map[string]interface{}{
		"transfer_no": transferDetail.TransferNo,
		"code":        transferDetail.ProductCode,
		"qty":         transferDetail.Qty,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transferDetail,
	})
}

func TransferDetailDelete(c fiber.Ctx) error {
	transferDetail := new(models.TransferDetail)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&transferDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transferDetail deleted",
	})
}
