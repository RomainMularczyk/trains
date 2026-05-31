package configTypes

import (
	"log/slog"
	"os"
)

type LoggingLevel int

const (
	Error LoggingLevel = iota
	Warn
	Info
	Debug
)

type Logging struct {
	Level LoggingLevel
}

type LoggingOverrides struct {
	Level *LoggingLevel
}

type StageLogger struct {
	Log    *slog.Logger
	stages []any
}

/*
Applies the given logging overrides to the logging configuration.
*/
func (l *Logging) Apply(o *LoggingOverrides) {
	if o == nil {
		return
	}
	if o.Level != nil {
		l.Level = *o.Level
	}
}

/*
Creates a new logger from the given configuration.
*/
func NewLogger(configLevel LoggingLevel) *StageLogger {
	var level slog.Level

	switch configLevel {
	case Error:
		level = slog.LevelError
	case Warn:
		level = slog.LevelWarn
	case Info:
		level = slog.LevelInfo
	case Debug:
		level = slog.LevelDebug
	default:
		level = slog.LevelInfo
	}

	base := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	)

	return &StageLogger{
		Log:    slog.New(base),
		stages: make([]any, 0),
	}
}

/*
Add a stage to the logger.
*/
func (s *StageLogger) Add(stage string, err error, result ...any) error {
	attrs := []slog.Attr{}

	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
		s.stages = append(s.stages, slog.GroupAttrs(stage, attrs...))

		return err
	}

	attrs = append(attrs, slog.Bool("success", true), slog.Any("result", result))
	s.stages = append(s.stages, slog.GroupAttrs(stage, attrs...))
	return nil
}

func (s *StageLogger) Debug(message string, args ...any) {
	s.Log.Debug(message)
}

func (s *StageLogger) Info(message string, args ...any) {
	s.Log.Info(message, slog.Group("stages", s.stages...))
}

func (s *StageLogger) Warn(message string, args ...any) {
	s.Log.Warn(message, slog.Group("stages", s.stages...))
}

func (s *StageLogger) Error(message string, args ...any) {
	s.Log.Error(message, slog.Group("stages", s.stages...))
}
