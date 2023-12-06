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
	server := app.NewServer(userService)
	server.Start(config.Get().Port)
}
