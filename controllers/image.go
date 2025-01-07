package controllers

import (
	"fmt"
	"inv_fiber/config"
	"inv_fiber/models"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ImageIndex(c fiber.Ctx) error {
	var images []models.ProductImage

	if res := config.DB.Debug().Find(&images); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "Image(s) not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": images,
	})
}

func ImagePost(c fiber.Ctx) error {
	image := new(models.ProductImage)

	if err := c.Bind().Body(image); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if image.ProductCode == "" ||
		image.FileName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "image upload cancelled"})
	}

	// Save file to the local file system
	pathmaster := filepath.Dir("../uploads/")
	filename := filepath.Base(file.Filename)
	filePath := filepath.Join(pathmaster, "images", filename)

	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	item := models.ProductImage{
		ID:          uuid.New(),
		ProductCode: image.ProductCode,
		FileName:    filePath,
	}
	config.DB.Debug().Create(&item)

	return c.JSON(fiber.Map{
		"data": image,
	})

}

func ImageShow(c fiber.Ctx) error {
	var images []models.ProductImage

	if result := config.DB.Debug().First(&images, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": images,
	})
}

func ImageDelete(c fiber.Ctx) error {
	var images models.ProductImage

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&images)

	var result string
	config.DB.Raw("select file_name from item_images Where id= ? ", id).Scan(&result)

	fmt.Println(result)
	if err := os.Remove(result); err != nil {
		log.Fatal(err)
	}

	config.DB.Delete(&images)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Image deleted",
	})

}
