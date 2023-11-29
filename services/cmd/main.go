package main

import (
	"com.clabs.com/services/v1/app"
	"github.com/gofiber/fiber/v2"
)

func main() {
	appFiber := fiber.New()
	app.LoadRoutes(appFiber)
	appFiber.Listen(":3000")
}
