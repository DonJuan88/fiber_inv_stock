package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func NotificationIndex(c fiber.Ctx) error {
	var notifications []*models.Notification

	if res := config.DB.Debug().Find(&notifications); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "notification not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": notifications,
	})
}

func NotificationShow(c fiber.Ctx) error {
	var notification []*models.Notification

	if result := config.DB.Debug().First(&notification, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "notification not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": notification,
	})

}

func NotificationCreate(c fiber.Ctx) error {
	notification := new(models.Notification)

	if err := c.Bind().Body(notification); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if notification.NotifId == "" ||
		notification.UserID == "" ||
		notification.NotifMessage == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(notification)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newNotification := models.Notification{

		ID:           uuid.New(),
		NotifId:      notification.NotifId,
		UserID:       notification.UserID,
		NotifMessage: notification.NotifMessage,
	}
	config.DB.Debug().Create(&newNotification)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": notification,
	})
}

func NotificationUpdate(c fiber.Ctx) error {
	notification := new(models.Notification)

	if err := c.Bind().Body(notification); err != nil {
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
	if notification.NotifId == "" ||
		notification.UserID == "" ||
		notification.NotifMessage == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.Notification{}).Where("id = ?", id).Updates(map[string]interface{}{
		"notif_id":      notification.NotifId,
		"user_id":       notification.UserID,
		"notif_message": notification.NotifMessage,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": notification,
	})
}

func NotificationDelete(c fiber.Ctx) error {
	notification := new(models.Notification)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&notification)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notification deleted",
	})
}

/* func notificationSearch(c fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */
