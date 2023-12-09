package config

type Config struct {
	Port string
	DB   string
}

var c *Config

func ParseEnv() {
	c = &Config{
		Port: "8080",
		DB:   "host=localhost port=5434 user=postgres password= dbname=checkmate sslmode=disable",
	}
}

func Get() *Config {
	return c
}
