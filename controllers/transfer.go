package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func TransferIndex(c fiber.Ctx) error {
	var transfers []*models.Transfer

	if res := config.DB.Debug().Find(&transfers); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "transfer not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transfers,
	})
}

func TransferShow(c fiber.Ctx) error {
	var transfer []*models.Transfer

	if result := config.DB.Debug().First(&transfer, c.Params("id=?")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "transfer not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transfer,
	})

}

func TransferCreate(c fiber.Ctx) error {
	transfer := new(models.Transfer)

	if err := c.Bind().Body(transfer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if transfer.TransferNo == "" ||
		transfer.TransferDate.GoString() == "" ||
		transfer.BranchOrigin == "" ||
		transfer.BranchDestiny == "" ||
		transfer.Notes == "" ||
		transfer.UserId == "" ||
		transfer.Cost < 0 ||
		transfer.Reference == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(transfer)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newtransfer := models.Transfer{

		ID:            uuid.New(),
		TransferNo:    transfer.TransferNo,
		TransferDate:  transfer.TransferDate,
		BranchOrigin:  transfer.BranchOrigin,
		BranchDestiny: transfer.BranchDestiny,
		Reference:     transfer.Reference,
		Notes:         transfer.Notes,
		UserId:        transfer.UserId,
		Cost:          transfer.Cost,
	}

	config.DB.Debug().Create(&newtransfer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transfer,
	})
}

func TransferUpdate(c fiber.Ctx) error {
	transfer := new(models.Transfer)

	if err := c.Bind().Body(transfer); err != nil {
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
	if transfer.TransferNo == "" ||
		transfer.TransferDate.GoString() == "" ||
		transfer.BranchOrigin == "" ||
		transfer.BranchDestiny == "" ||
		transfer.Notes == "" ||
		transfer.UserId == "" ||
		transfer.Cost < 0 ||
		transfer.Reference == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Transfer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"transfer_no":    transfer.TransferNo,
		"transfer_date":  transfer.TransferDate,
		"branch_origin":  transfer.BranchOrigin,
		"branch_destiny": transfer.BranchDestiny,
		"reference":      transfer.Reference,
		"notes":          transfer.Notes,
		"user_id":        transfer.UserId,
		"cost":           transfer.Cost})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": transfer,
	})
}

func TransferDelete(c fiber.Ctx) error {
	transfer := new(models.Transfer)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&transfer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "transfer deleted",
	})
}
