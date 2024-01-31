package middleware

import (
	"fmt"
	"strconv"

	"github.com/devshahorair/fiber/utils"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	user, err := utils.GetUserDataFromJWT(c.Get("Authorization"))
	if err != nil {
		return c.JSON(fiber.Map{
			"error": "login and send authorization",
		})
	}

	c.Set("userEmail", user.Email)
	c.Set("userId", strconv.Itoa(int(user.Id)))
	fmt.Println(user.Email)
	return c.Next()
}
