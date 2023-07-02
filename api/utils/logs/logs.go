package logs

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var (
	logger *zerolog.Logger
	once   sync.Once
)

// GetLogger returns a singleton logger instance.
func GetLogger() *zerolog.Logger {
	once.Do(func() {
		logger = newLogger()
	})
	return logger
}

// newLogger creates a new Logger instance.
func newLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &l
}

// Log records a message using the logger.
func Log(message string) {
	GetLogger().Info().Msg(message)
}
