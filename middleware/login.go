package middleware

import (
	"inv_fiber/config"
	"inv_fiber/models"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func UserLogin(c fiber.Ctx) error {
	var loginDetails models.Login

	if loginDetails.Email == "" ||
		loginDetails.Password == "" {

		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	if err := c.Bind().Body(&loginDetails); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var account models.User
	if err := config.DB.Where("email = ? ", loginDetails.Email).First(&account).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "nvalid username or password",
		})

	}

	// Compare password
	comparePasswod := account.Password //+ string(config.ENV.TOKEN_LOGIN)
	if err := bcrypt.CompareHashAndPassword([]byte(comparePasswod), []byte(loginDetails.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "nvalid username or password",
		})
	}

	//	fmt.Println("step 3")
	generateToken := jwt.New(jwt.SigningMethodHS256)

	claims := generateToken.Claims.(jwt.MapClaims)
	claims["identity"] = account
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := generateToken.SignedString([]byte(config.ENV.TOKEN_LOGIN)) // ---ASLI---

	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//	fmt.Println("step 4")

	return c.JSON(fiber.Map{"message": "Success login", "data": token})

}

func Accessible(c fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

/*
func Login(c fiber.Ctx) error {
 var req authRequest
 if err := c.Bind().Body(&req); err != nil {
  return c.Status(400).JSON(fiber.Map{
   "message": err.Error(),
  })
 }
 var user model.User
 res := database.DB.Where("email = ?", req.Email).First(&user)
 if res.Error != nil {
  return c.Status(400).JSON(fiber.Map{
   "message": "user not found",
  })
 }
 if !utils.ComparePassword(user.PasswordHash, req.Password) {
  return c.Status(400).JSON(fiber.Map{
   "message": "incorrect password",
  })
 }

 token, err := utils.GenerateToken(user.ID)
 if err != nil {
  return c.Status(500).JSON(fiber.Map{
   "message": err.Error(),
  })
 }
 return c.JSON(fiber.Map{
  "token": token,
 })
}





import (
 "os"

 "github.com/gofiber/fiber/v3"

 jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected(c fiber.Ctx) error {
 return jwtware.New(jwtware.Config{
  SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
  ContextKey: "jwt",
  ErrorHandler: func(c fiber.Ctx, err error) error {
   // Return status 401 and failed authentication error.
   return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
    "error": true,
    "msg":   err.Error(),
   })
  },
 })(c)
}

*/
