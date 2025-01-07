package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func BranchStockIndex(c fiber.Ctx) error {
	var branchStocks []*models.BranchStockPrice

	if res := config.DB.Debug().Find(&branchStocks); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "BranchStock not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branchStocks,
	})
}

func BranchStockShowAll(c fiber.Ctx) error {
	var branchStocks []*models.BranchStockPrice

	if result := config.DB.Debug().First(&branchStocks, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BranchStock not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branchStocks,
	})

}

func BranchStockShowStore(c fiber.Ctx) error {
	var branchStocks []*models.BranchStockPrice

	if result := config.DB.Debug().First(&branchStocks, "branch_code=?", c.Params("branch_code")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "BranchStock not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branchStocks,
	})

}

func BranchStockCreate(c fiber.Ctx) error {
	branchStock := new(models.BranchStockPrice)

	if err := c.Bind().Body(branchStock); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if branchStock.BranchCode == "" ||
		branchStock.ProductCode == "" ||
		branchStock.Barcode1 == "" ||
		branchStock.Barcode2 == "" ||
		branchStock.BasePrice <= 0 ||
		branchStock.SalePrice <= 0 ||
		branchStock.Stock <= 0 ||
		branchStock.MinStock <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if BranchStock exists

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(branchStock)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newBranchStock := models.BranchStockPrice{

		ID:          uuid.New(),
		BranchCode:  branchStock.BranchCode,
		ProductCode: branchStock.ProductCode,
		Barcode1:    branchStock.Barcode1,
		Barcode2:    branchStock.Barcode2,
		BasePrice:   branchStock.BasePrice,
		SalePrice:   branchStock.SalePrice,
		Stock:       branchStock.Stock,
		MinStock:    branchStock.MinStock,
	}
	//
	// hashPassword, err := utils.HashPassword(user.Password)
	// if err != nil {
	//  return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//   "message": "status internal server error",
	//  })
	// }
	//
	//newUser.Password = hashPassword

	config.DB.Debug().Create(&newBranchStock)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Created new BranchStock",
	})
}

func BranchStockUpdate(c fiber.Ctx) error {
	branchStock := new(models.BranchStockPrice)

	if err := c.Bind().Body(branchStock); err != nil {
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

	config.DB.Debug().Model(&models.BranchStockPrice{}).Where("id = ?", id).Updates(map[string]interface{}{
		"branch_code": branchStock.BranchCode,
		"code":        branchStock.ProductCode,
		"barcode1":    branchStock.Barcode1,
		"barcode2":    branchStock.Barcode2,
		"baseprice":   branchStock.BasePrice,
		"saleprice":   branchStock.SalePrice,
		"stock":       branchStock.Stock,
		"min_stock":   branchStock.MinStock,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succes update detail stock",
	})
}

/* func BranchStockDelete(c fiber.Ctx) error {
	BranchStocks := new(models.BranchStockPrice)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&BranchStocks)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete BranchStock successfully",
	})
} */

/* func BranchStockSearch(c fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */
