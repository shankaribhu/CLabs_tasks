package controller

import (
	"fmt"

	"com.clabs.com/services/v1/app/common"
	"github.com/gofiber/fiber/v2"
)

type ReceiverController struct{}

func NewReceiverController() *ReceiverController {
	p := ReceiverController{}
	return (&p)
}

func (p *ReceiverController) SubmitReq(ctx *fiber.Ctx) error {
	receiverData := map[string]interface{}{}
	err := ctx.BodyParser(&receiverData)
	fmt.Println("Err", err)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	convertData := common.ConvertData(receiverData)
	return ctx.Status(fiber.StatusOK).JSON(convertData)
}
