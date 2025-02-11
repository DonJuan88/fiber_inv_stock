package controllers

import (
	"inv_fiber/config"
	"inv_fiber/helper"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func CustomerIndex(c fiber.Ctx) error {
	var customers []*models.Customer

	if res := config.DB.Debug().Find(&customers); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Customer not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": customers,
	})
}
func CustomerShow(c fiber.Ctx) error {
	var customer []*models.Customer

	if result := config.DB.Debug().First(&customer, c.Params("id=?")); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Customer not found",
		})
	}
	//if result := config.DB.Debug().First(&customer, c.Params("id")); result.Error != nil {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": "customers not found",
	//	})
	//}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": customer,
	})

}

func CustomerCreate(c fiber.Ctx) error {
	customer := new(models.Customer)

	if err := c.Bind().Body(customer); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if customer.CustomerID == "" ||
		customer.CustomerName == "" ||
		customer.ContactPerson == "" ||
		customer.Email == "" ||
		customer.Phone == "" ||
		customer.Address == "" ||
		customer.City == "" ||
		customer.State == "" ||
		customer.PostalCode == "" ||
		customer.Country == "" {

		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Check if customer exists
	exists, err := helper.CheckCustomerExists(config.DB, customer.CustomerID)
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
	errValidate := validate.Struct(customer)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newcustomer := models.Customer{

		ID:            uuid.New(),
		CustomerID:    customer.CustomerID,
		CustomerName:  customer.CustomerName,
		ContactPerson: customer.ContactPerson,
		Email:         customer.Email,
		Phone:         customer.Phone,
		Address:       customer.Address,
		City:          customer.City,
		State:         customer.State,
		PostalCode:    customer.PostalCode,
		Country:       customer.Country,
	}
	config.DB.Debug().Create(&newcustomer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": customer,
	})
}

func CustomerUpdate(c fiber.Ctx) error {
	customer := new(models.Customer)

	if err := c.Bind().Body(customer); err != nil {
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
	if customer.CustomerID == "" ||
		customer.CustomerName == "" ||
		customer.ContactPerson == "" ||
		customer.PostalCode == "" ||
		customer.Email == "" ||
		customer.Phone == "" ||
		customer.Address == "" ||
		customer.City == "" ||
		customer.State == "" ||
		customer.Country == "" {

		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Customer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"code":           customer.CustomerID,
		"name":           customer.CustomerName,
		"contact_person": customer.ContactPerson,
		"email":          customer.Email,
		"phone":          customer.Phone,
		"address":        customer.Address,
		"city":           customer.City,
		"state":          customer.State,
		"postalcode":     customer.PostalCode,
		"country":        customer.Country,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": customer,
	})
}

func CustomerDelete(c fiber.Ctx) error {
	customer := new(models.Customer)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&customer)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer deleted",
	})
}

/* func customerearch(c fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */
