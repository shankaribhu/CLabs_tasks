package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	controllerReceiver := NewReceiverController()
	router.Post("/submitreq", controllerReceiver.SubmitReq)
	controllerWorker := NewWorkerController()
	router.Post("/submitworker", controllerWorker.SubmitWorker)
	controllerChannel := NewChannelController()
	router.Post("/submitchannel", controllerChannel.SubmitChannel)
}
