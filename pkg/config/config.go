package config

import "os"

type Config struct {
	Port string
	DB   string
}

var c *Config

func ParseEnv() {
	c = &Config{
		Port: os.Getenv("PORT"),
		DB:   os.Getenv("DB"),
	}
}

func Get() *Config {
	return c
}
