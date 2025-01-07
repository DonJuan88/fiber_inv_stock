package controllers

import (
	"inv_fiber/config"
	"inv_fiber/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/* func SaleDetailIndex(c *fiber.Ctx) error {
	var saleDetails []*models.SaleDetailetail

	if res := config.DB.Debug().Find(&saleDetails); res.Error != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "saleDetail not found",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": saleDetails,
	})
}
*/
func SaleDetailShow(c *fiber.Ctx) error {
	var saleDetail []*models.SaleDetail

	if result := config.DB.Debug().First(&saleDetail, c.Params("id")); result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "sale not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": saleDetail,
	})

}

func SaleDetailCreate(c *fiber.Ctx) error {
	saleDetail := new(models.SaleDetail)

	if err := c.BodyParser(saleDetail); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Field incomplete",
		})
	}

	// Validate required fields
	if 
	saleDetail.SaleNo    == "" || 
	saleDetail.SaleDate.GoString() == "" ||
	saleDetail.ItemCode == "" ||
	saleDetail.Qty      <0 ||
	saleDetail.Price    <0 ||
	saleDetail.Discount <0 ||
	saleDetail.SalePrice<0   {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	
	// Validation
	validate := validator.New()
	errValidate := validate.Struct(saleDetail)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
		})
	}

	newsaleDetail := models.SaleDetail{

	ID:                 uuid.New(),
	SaleNo    : saleDetail.SaleNo    ,
	SaleDate  : saleDetail.SaleDate  ,
	ItemCode      : saleDetail.ItemCode      ,
	Qty           : saleDetail.Qty           ,
	Price     : saleDetail.Price     ,
	Discount      : saleDetail.Discount     ,
	SalePrice : saleDetail.SalePrice , }
	config.DB.Debug().Create(&newsaleDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": saleDetail,
	})
}

func SaleDetailUpdate(c *fiber.Ctx) error {
	saleDetail := new(models.SaleDetail)

	if err := c.BodyParser(saleDetail); err != nil {
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
	if 
	
	saleDetail.SaleNo    == "" || 
	saleDetail.SaleDate.GoString()  == "" || 
	saleDetail.ItemCode      == "" || 
	saleDetail.Qty          <0  || 
	saleDetail.Price     <0 ||
	saleDetail.Discount      <0 ||
	saleDetail.SalePrice <0  {
		return c.Status(400).JSON(fiber.Map{"error": "Complete the fields"})
	}

	config.DB.Debug().Model(&models.SaleDetail{}).Where("id = ?", id).Updates(map[string]interface{}{
	  "sale_no" :saleDetail.SaleNo,      
	  "sale_date":saleDetail.SaleDate,    
	  "code"        :saleDetail.ItemCode,        
	  "qty"         :saleDetail.Qty,         
 	  "baseprice"   :saleDetail.Price,     
	  "discount"    :saleDetail.Discount,      
	 "saleprice" :saleDetail.SalePrice,    })
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": saleDetail,
	})
}

func SaleDetailDelete(c *fiber.Ctx) error {
	saleDetail := new(models.SaleDetail)

	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	config.DB.Debug().Where("id = ?", id).Delete(&saleDetail)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "saleDetail deleted",
	})
}
