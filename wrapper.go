// Package restyzerolog provides a wrapper for [zerolog.Logger] to be used with resty
// See:
//   - https://resty.dev
//   - https://pkg.go.dev/github.com/go-resty/resty/v3#Logger
package restyzerolog

import (
	"github.com/rs/zerolog"
)

// Logger is a wrapper for [zerolog.Logger] to be used as logger for resty.
// Contains an instance of [zerolog.Logger], which is used to log messages.
// Not exported, because it's not necessary to use it directly.
type Logger struct {
	logger zerolog.Logger
}

// New creates a new instance of [Logger] with provided [zerolog.Logger].
//
// Example of wrapping the default global zerolog logger:
//
//	client := resty.New()
//	client.SetLogger(restyzerolog.New(log.Logger))
//
// See:
//
//   - https://pkg.go.dev/github.com/rs/zerolog/log#pkg-variables
func New(logger zerolog.Logger) *Logger {
	return &Logger{
		logger: logger,
	}
}

func (r *Logger) Errorf(format string, v ...any) {
	r.logger.Error().Msgf(format, v...)
}

func (r *Logger) Warnf(format string, v ...any) {
	r.logger.Warn().Msgf(format, v...)
}

func (r *Logger) Debugf(format string, v ...any) {
	r.logger.Debug().Msgf(format, v...)
}
