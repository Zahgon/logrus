package slog

import (
	"log/slog"

	"github.com/sirupsen/logrus"
)

const (
	slogLevelTrace = slog.LevelDebug - 4
	slogLevelDebug = slog.LevelDebug
	slogLevelInfo  = slog.LevelInfo
	slogLevelWarn  = slog.LevelWarn
	slogLevelError = slog.LevelError
	slogLevelFatal = slog.LevelError + 2
	slogLevelPanic = slog.LevelError + 4
)

type Leveler interface {
	Level() logrus.Level
}

type Level logrus.Level

func (l Level) Level() slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }

var _ slog.Leveler = Level(0)

type SlogLevel slog.Level

func (l SlogLevel) Level() logrus.Level { _ = "STUB: not implemented"; return *new(logrus.Level) }

var _ Leveler = SlogLevel(0)

func toLogrusLevel(level slog.Level) logrus.Level {
	_ = "STUB: not implemented"
	return *new(logrus.Level)
}

func toSlogLevel(level logrus.Level) slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }
