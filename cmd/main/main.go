package main

import (
	"check-mate/internal/app"
	"check-mate/internal/service"
	"check-mate/internal/store"
	"check-mate/pkg/config"
	"check-mate/pkg/db"
	"log"
)

func init() {
	config.ParseEnv()
}

func main() {
	db, err := db.Connect(config.Get().DB)
	if err != nil {
		log.Fatalf("failed to connect db: %s", err)
	}

	s := store.NewStore(db)

	userService := service.NewUserService(s)
	migrationService := service.NewMigrationService(s)
	messagesService := service.NewMessagesService(s)
	server := app.NewServer(userService, messagesService, migrationService)
	server.Route()
	err = server.Start(config.Get().Port)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
