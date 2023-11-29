package app

import (
	router_controller "com.clabs.com/services/v1/app/controllers/receiver"
	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(router *fiber.App) {
	routerRequest := router.Group("/v1")
	routerRequest.Route("/", router_controller.Router)
}
