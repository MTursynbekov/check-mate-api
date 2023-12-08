package app

import (
	"check-mate/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app         *fiber.App
	userService service.UsersService
	messagesService service.MessagesService
}

func NewServer(userService service.UsersService, messagesService service.MessagesService) *Server {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	s := &Server{
		app:         app,
		userService: userService,
		messagesService: messagesService,
	}

	return s
}

func (s *Server) Start(port string) error {
	return s.app.Listen(":" + port)
}
