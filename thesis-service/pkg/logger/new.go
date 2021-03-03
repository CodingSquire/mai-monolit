package logger

import (
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"io"
	"os"

	"github.com/rs/zerolog"


)

func NewLogger(cfg *Config) *api.Logger {
	var output io.Writer = os.Stdout
	if cfg.Pretty {
		output = zerolog.ConsoleWriter{Out: output}
	}
	logger := zerolog.New(output).With().Logger()

	if cfg.Timestamp {
		logger = logger.With().Timestamp().Logger()
	}
	if cfg.Caller {
		logger = logger.With().Caller().Logger()
	}

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		logger.Warn().Err(err).Str("level", cfg.Level).Msg("Cannot parse a logging level")
	} else {
		logger = logger.Level(level)
	}

	return &logger
}
