package app

import (
	"check-mate/internal/model"
	"encoding/json"
	"log"
	"strconv"

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

func (S *Server)GetMessages(c *fiber.Ctx) error{
	chatId := c.Params("chatId")
	id, _ := strconv.Atoi(chatId)

	messages, err := S.messagesService.GetMessages(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(messages)
}

func (S *Server)CreateChat(c *fiber.Ctx) error{
	chat := new(model.Chat)
	reqBody := c.Request().Body()

	err := json.Unmarshal(reqBody, &chat)
	if err != nil{
		log.Println("error while unmarshalling request")
		return c.Status(400).SendString("invalid request")
	}

	err = S.messagesService.CreateChat(chat)
	if err != nil{
		log.Println("error while creating chat")
		return c.Status(500).SendString("error while creating chat")
	}

	return c.JSON(chat)
}