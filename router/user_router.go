package router

import (
	"errors"
	"fmt"

	"github.com/devshahorair/fiber/db"
	"github.com/devshahorair/fiber/models"
	"github.com/devshahorair/fiber/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	type Post struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Pass  string `json:"pass" validate:"required"`
	}
	body := new(Post)
	c.BodyParser(body)

	if err := validate.Struct(body); err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": "name, email, pass are required filds.",
		})
	}
	hashedPass := utils.HashPassword(body.Pass)
	result := db.DB.Create(&models.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: hashedPass,
	})

	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {

		return c.JSON(fiber.Map{
			"message": body.Email + " this email alrady registered.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user register success",
	})
}

func Login(c *fiber.Ctx) error {
	type Post struct {
		Email string `json:"email" validate:"required,email"`
		Pass  string `json:"pass" validate:"required"`
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	body := new(Post)
	c.BodyParser(body)
	if err := validate.Struct(body); err != nil {
		return c.JSON(fiber.Map{
			"error": "email, pass required filds.",
		})
	}

	var user models.Users

	if err := db.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": "this email not register.",
		})
	}
	if utils.CheckPasswordHash(body.Pass, user.Password) {
		return c.JSON(fiber.Map{
			"error": "wrong password.",
		})
	}

	return c.JSON(fiber.Map{
		"jwt_token": utils.GetJwt(user.ID, user.Email),
	})
}

func Me(c *fiber.Ctx) error {
	var user models.Users
	fmt.Println("from merouter = " + c.Get("userEmail"))
	db.DB.First(&user, c.Get("userId"))

	return c.JSON(fiber.Map{
		"name":  user.Name,
		"email": user.Email,
		"id":    user.ID,
	})
}
