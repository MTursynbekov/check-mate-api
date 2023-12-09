package app

import (
	"check-mate/internal/service"
	"check-mate/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

type Server struct {
	app             *fiber.App
	userService     service.UsersService
	migration       service.MigrationService
	contactService  service.ContactsService
	messagesService service.MessagesService
}

func NewServer(userService service.UsersService, migrationService service.MigrationService, contactService service.ContactsService, messagesService service.MessagesService) *Server {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Use(logger.New())
	app.Use("/api/", jwtware.New(jwtware.Config{
		SigningKey: []byte(config.SECRET),
	}))

	s := &Server{
		app:             app,
		userService:     userService,
		messagesService: messagesService,
		migration:       migrationService,
		contactService:  contactService,
	}

	s.Route()

	return s
}

func (s *Server) Start(port string) error {
	s.migration.Migrate()
	return s.app.Listen(":" + port)
}
