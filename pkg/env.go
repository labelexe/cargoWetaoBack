package pkg

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func EnvLoadInit() {
	err := godotenv.Load()
	if err != nil {
		log.Err(fmt.Errorf("error loading .env file - %s", err))
	}
}
