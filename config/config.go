package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

// Schema struct
type Schema struct {
	PancakeRouter int
	TokenPath     []string
}

// Config global parameter config
var Config *Schema

func LoadConfig() *Schema {
	_ = godotenv.Load(".env")

	pancakeRouter := os.Getenv("PANCAKE_ROUTER")
	tokenPath := strings.Split(os.Getenv("PANCAKE_ROUTER"))

	Config = &Schema{
		PancakeRouter: pancakeRouter,
		TokenPath:     tokenPath,
	}

	return Config
}
