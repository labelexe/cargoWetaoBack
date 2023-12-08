package pkg

import (
	"github.com/rs/zerolog"
	"os"
)

func LoggerInit() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info()
}
