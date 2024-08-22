package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Gateway struct {
	UserHost string
	Userport string
}

func GatewayConfig() (*Gateway, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Gateway{
		UserHost: os.Getenv("USER_HOST"),
		Userport: os.Getenv("USER_PORT"),
	}, nil
}
