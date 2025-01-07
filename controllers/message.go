package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func MessageShow(c *fiber.Ctx) error {
	var messages []*models.Message

	if result := config.DB.Debug().First(&messages, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data": "message not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": messages,
	})

}

func MessageCreate(c *fiber.Ctx) error {
	message := new(models.Message)

	if err := c.BodyParser(message); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 
	message.Name == "" ||
	message.EmailPhone          == "" ||
	message.MyMessage == "" ||
	!message.ReadingStatus {
		
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}


	// Validation
	validate := validator.New()
	errValidate := validate.Struct(message)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"data": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newmessage := models.Message{

		ID: uuid.New(),
	Name          : message.Name,
	EmailPhone    : message.EmailPhone,
	MyMessage     : message.MyMessage,
	ReadingStatus :message.ReadingStatus,
	}
	config.DB.Debug().Create(&newmessage)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": message,
	})
}


func MessageDelete(c *fiber.Ctx) error {
	message := new(models.Message)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&message)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "Message deleted",
	})
}

/* func messageearch(c *fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */