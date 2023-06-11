package logs

import (
	"github.com/rs/zerolog"
	"os"
)

var (
	logger *zerolog.Logger
)

// GetLogger devuelve una instancia única del logger.
func GetLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if logger == nil {
		logger = newLogger()
	}

	return logger
}

// newLogger crea una nueva instancia de Logger.
func newLogger() *zerolog.Logger {
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &l
}

// Log registra un mensaje utilizando el logger.
func Log(message string) {
	GetLogger().Info().Msg(message)
}
