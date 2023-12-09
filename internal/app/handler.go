package app

import (
	"check-mate/internal/model"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) CreateMessage(c *fiber.Ctx) error {
	msg := new(model.Message)
	reqBody := c.Request().Body()

	err := json.Unmarshal(reqBody, &msg)
	if err != nil {
		log.Println("error while unmarshalling message")
		return c.Status(400).SendString("invalid request")
	}

	err = s.messagesService.CreateMessage(msg)
	if err != nil {
		log.Println("error while storing message")
		return c.Status(500).SendString("error while storing message")
	}

	return c.Status(200).SendString("Message sent!")
}

func (s *Server) GetMessages(c *fiber.Ctx) error {
	chatId := c.Params("chatId")
	id, _ := strconv.Atoi(chatId)

	messages, err := s.messagesService.GetMessages(id)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(messages)
}

func (s *Server) CreateChat(c *fiber.Ctx) error {
	chat := new(model.Chat)
	reqBody := c.Request().Body()

	err := json.Unmarshal(reqBody, &chat)
	if err != nil {
		log.Println("error while unmarshalling request")
		return c.Status(400).SendString("invalid request")
	}

	err = s.messagesService.CreateChat(chat)
	if err != nil {
		log.Println("error while creating chat")
		return c.Status(500).SendString("error while creating chat")
	}

	return c.JSON(chat)
}

func (s *Server) CreateContact(c *fiber.Ctx) error {
	contact := new(model.Contact)
	reqBody := c.Request().Body()

	err := json.Unmarshal(reqBody, &contact)
	if err != nil {
		log.Println("error while unmarshalling request")
		return c.Status(400).SendString("invalid request")
	}

	err = s.contactService.CreateContact(contact)
	if err != nil {
		log.Println("error while creating chat")
		return c.Status(500).SendString("error while creating chat")
	}

	return c.JSON(contact)
}
