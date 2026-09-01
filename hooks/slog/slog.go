package slog

import (
	"context"
	"log/slog"

	"github.com/sirupsen/logrus"
)

type logrusLoggerContextKey struct{}

func withLogrusLogger(ctx context.Context, logger *logrus.Logger) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func hasLogrusLogger(ctx context.Context, logger *logrus.Logger) bool {
	_ = "STUB: not implemented"
	return false
}

type HookOptions struct {
	LevelMapper func(logrus.Level) slog.Level
}

type Hook struct {
	logger *slog.Logger
	opts   HookOptions
}

var _ logrus.Hook = (*Hook)(nil)

func NewHook(logger *slog.Logger, opts *HookOptions) *Hook { _ = "STUB: not implemented"; return nil }

func (h *Hook) toSlogLevel(level logrus.Level) slog.Level {
	_ = "STUB: not implemented"
	return *new(slog.Level)
}

func (h *Hook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (h *Hook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }
