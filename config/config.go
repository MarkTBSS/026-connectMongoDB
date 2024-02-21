package config

import (
	"os"

	"github.com/joho/godotenv"
)

// type aliasing
type (
	Config struct {
		Db Db
	}

	Db struct {
		Url string
	}
)

func LoadConfig(path string) Config {
	godotenv.Load(path)
	return Config{
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
	}
}
