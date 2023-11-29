package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type WorkerController struct {
	requests chan map[string]interface{}
}

func NewWorkerController() *WorkerController {
	p := WorkerController{
		requests: make(chan map[string]interface{}, 100),
	}
	go p.worker()
	return (&p)
}

func (p *WorkerController) worker() {
	for receiverData := range p.requests {
		fmt.Println("Wroker processing data", receiverData)
	}
}

func (p *WorkerController) SubmitWorker(ctx *fiber.Ctx) error {
	receiverData := map[string]interface{}{}
	err := ctx.BodyParser(&receiverData)
	fmt.Println("Err", err)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	p.requests <- receiverData
	return ctx.Status(fiber.StatusOK).JSON("Worker Data Processed Successfully!")
}
