package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CompanyShow(c *fiber.Ctx) error {
	var company []*models.Company

	if result := config.DB.Debug().First(&company, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "company not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": company,
	})

}

func CompanyCreate(c *fiber.Ctx) error {
	company := new(models.Company)

	if err := c.BodyParser(company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 
	company.Name  == "" ||
	company.ContactPerson       == "" ||
	company.ContactPersonPhone  == "" ||
	company.Email               == "" ||
	company.Phone               == "" ||
	company.Address             == "" ||
	company.City               == "" ||
	company.State               == "" ||
	company.PostalCode         == "" ||
	company.Country             == ""  {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(company)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newcompany := models.Company{

		ID:          uuid.New(),
	Name       : company.Name,
	ContactPerson:company.ContactPerson,
	ContactPersonPhone: company.ContactPersonPhone,
	Email:company.Email,
	Phone: company.Phone,
	Address:company.Address,
	City:company.City,
	State:company.State,
	PostalCode:company.PostalCode,
	Country: company.Country,
	}
	config.DB.Debug().Create(&newcompany)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": company,
	})
}

func CompanyUpdate(c *fiber.Ctx) error {
	company := new(models.Company)

	if err := c.BodyParser(company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Model(&models.Company{}).Where("id = ?", id).Updates(map[string]interface{}{
	"name" : company.Name,
	"cp":company.ContactPerson,
	"cp_phone": company.ContactPersonPhone,
	"email":company.Email,
	"phone": company.Phone,
	"address":company.Address,
	"city":company.City,
	"state":company.State,
	"postalcode":company.PostalCode,
	"country": company.Country,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": company,
	})
}

func CompanyDelete(c *fiber.Ctx) error {
	company := new(models.Company)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&company)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Company deleted",
	})
}

