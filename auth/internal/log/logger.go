package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

var (
	globalLogger *zerolog.Logger // nolint: gochecknoglobals

	ErrLogLevelParse = errors.New("log level parse error")
)

func SetupLogger(appName string) error {
	logLevel, err := zerolog.ParseLevel("INFO")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLogLevelParse, err.Error())
	}

	zerolog.SetGlobalLevel(logLevel)

	l := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("app", appName).
		Logger()

	globalLogger = &l

	return nil
}

func Default() *zerolog.Logger {
	return globalLogger
}
