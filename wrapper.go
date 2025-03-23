package restyzerolog

import "github.com/rs/zerolog"

type Logger struct {
	logger zerolog.Logger
}

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
