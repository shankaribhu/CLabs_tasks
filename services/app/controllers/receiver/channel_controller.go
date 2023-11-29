package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"

	"com.clabs.com/services/v1/app/common"
)

type ChannelController struct {
	requests chan map[string]interface{}
}

func NewChannelController() *ChannelController {
	p := ChannelController{
		requests: make(chan map[string]interface{}, 100),
	}
	go p.worker()
	return (&p)
}

func (p *ChannelController) SubmitChannel(ctx *fiber.Ctx) error {
	receiverData := map[string]interface{}{}
	err := ctx.BodyParser(&receiverData)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	p.requests <- receiverData
	convertData := common.ConvertData(receiverData)
	return ctx.Status(fiber.StatusOK).JSON(convertData)
}

func (p *ChannelController) worker() {
	for receiverData := range p.requests {
		convertData := common.ConvertData(receiverData)
		sendToWebhook(convertData)
	}
}

func sendToWebhook(data map[string]interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error in Marshalling", err)
		return
	}

	webhookURL := os.Getenv("WEBHOOKURL") // Replace with your actual webhook URL
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending data to webhook:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Data sent to webhook. Status:", resp.Status)
}
