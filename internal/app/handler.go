package app

import (
	"check-mate/internal/model"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (S *Server)CreateMessage(c *fiber.Ctx) error{
	msg := new(model.Message)
	reqBody := c.Request().Body()

	err := json.Unmarshal(reqBody, &msg)
	if err != nil{
		log.Println("error while unmarshalling message")
		return c.Status(400).SendString("invalid request")
	}

	err = S.messagesService.CreateMessage(msg)
	if err != nil{
		log.Println("error while storing message")
		return c.Status(500).SendString("error while storing message")
	}

	return c.Status(200).SendString("Message sent!")
}