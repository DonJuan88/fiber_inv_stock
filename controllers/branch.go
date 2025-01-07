package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func BranchIndex(c fiber.Ctx) error {
	var branchs []*models.Branch

	if res := config.DB.Debug().Find(&branchs); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Branch not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branchs,
	})
}

func BranchShow(c fiber.Ctx) error {
	var branch []*models.Branch

	if result := config.DB.Debug().First(&branch, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Branch not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branch,
	})

}

func BranchCreate(c fiber.Ctx) error {
	branch := new(models.Branch)

	if err := c.Bind().Body(branch); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if branch.BranchCode == "" ||
		branch.BranchName == "" ||
		branch.BranchAddress == "" ||
		branch.ContactPerson == "" ||
		branch.ContactPersonPhone == "" ||
		branch.Phone == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if Branch exists
	exists, err := helper.CheckBranchExists(config.DB, branch.BranchCode)
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
	errValidate := validate.Struct(branch)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newBranch := models.Branch{

		ID:                 uuid.New(),
		BranchCode:         branch.BranchCode,
		BranchName:         branch.BranchName,
		BranchAddress:      branch.BranchAddress,
		ContactPerson:      branch.ContactPerson,
		ContactPersonPhone: branch.ContactPersonPhone,
		Phone:              branch.Phone,
		Active:             true,
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

	config.DB.Debug().Create(&newBranch)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Created new Branch",
	})
}

func BranchUpdate(c fiber.Ctx) error {
	branch := new(models.Branch)

	if err := c.Bind().Body(branch); err != nil {
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
	if branch.BranchCode == "" ||
		branch.BranchName == "" ||
		branch.BranchAddress == "" ||
		branch.ContactPerson == "" ||
		branch.ContactPersonPhone == "" ||
		branch.Phone == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Branch{}).Where("id = ?", id).Updates(map[string]interface{}{
		"branch_code":          branch.BranchCode,
		"branch_name":          branch.BranchName,
		"branch_address":       branch.BranchAddress,
		"contact_person":       branch.ContactPerson,
		"contact_person_phone": branch.ContactPersonPhone,
		"phone":                branch.Phone,
		"active":               branch.Active,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": branch,
	})
}

func BranchDelete(c fiber.Ctx) error {
	branch := new(models.Branch)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&branch)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "delete branch successfully",
	})
}

/* func BranchSearch(c fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */
