package controllers

import (
	"inv_fiber/config"

	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func UserIndex(c *fiber.Ctx) error {
	var users []*models.User

	if res := config.DB.Debug().Find(&users); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}
func UserShow(c *fiber.Ctx) error {
	var user []*models.User

	if result := config.DB.Debug().First(&user, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "users not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})

}

func UserCreate(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 
	user.FirstName == "" ||
	user.LastName  == "" ||
	user.Email     == "" ||
	user.Password  == "" ||
	!user.IsAdmin    {
		
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}


	

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValidate.Error(),
		})
	}

	newuser := models.User{

		ID: uuid.New(),
		
	FirstName : user.FirstName,
	LastName  : user.LastName ,
	Email     : user.Email    ,
	Password  : user.Password ,
	IsAdmin   : user.IsAdmin  ,	}
	config.DB.Debug().Create(&newuser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
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
	if 
	
	user.FirstName == "" ||
	user.LastName  == "" ||
	user.Email     == "" ||
	user.Password  == "" ||
	!user.IsAdmin    {
		
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.User{}).Where("id = ?", id).Updates(map[string]interface{}{
	 "firstname" :user.FirstName,    
	 "lastname"  :user.LastName ,   
	 "email"     :user.Email    , 
	 "password"  :user.Password , 
	 "admin"     :user.IsAdmin  ,  
	 "active"    :user.Active   ,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func UserDelete(c *fiber.Ctx) error {
	user := new(models.User)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted",
	})
}

/* func Userearch(c *fiber.Ctx) error {
	query := c.Query("q")

	// Respond with the query parameter
	return c.JSON(fiber.Map{"query": query})
} */