package log

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/ulibaysya/krona/internal/config"
)

type Logger struct {
	Logger zerolog.Logger
	writer io.Writer
}

func New(cfg config.Log) (Logger, error) {
	const f = "github.com/ulibaysya/krona/internal/log.New"

	logger := Logger{}
	switch cfg.Path {
	case "stdout":
		logger.writer = os.Stdout
	case "stderr", "":
		logger.writer = os.Stderr
	default:
		var err error
		logger.writer, err = os.OpenFile(cfg.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return Logger{}, fmt.Errorf("%s: %w", f, err)
		}
	}

	logger.Logger = zerolog.New(logger.writer)

	return logger, nil
}

func (l Logger) Close() error {
	file, ok := l.writer.(io.Closer)
	if !ok {
		return errors.New("writer is not closure")
	}

	return file.Close()
}
