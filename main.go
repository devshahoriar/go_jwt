package main

import (
	"github.com/devshahorair/fiber/db"
	"github.com/devshahorair/fiber/middleware"
	"github.com/devshahorair/fiber/router"
	"github.com/gofiber/fiber/v2"
)

// to run use this command:
// CompileDaemon -command="./fiber.exe"

func init() {
	db.Db_Connect()
	// db.DB.AutoMigrate(&models.Users{})
}

func main() {
	app := fiber.New()
	app.Post("/register", router.Register)
	app.Post("/login", router.Login)
	app.Get("/me", middleware.CheckAuth, router.Me)
	app.Listen(":3000")
}
