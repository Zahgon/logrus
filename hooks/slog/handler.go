package slog

import (
	"context"
	"log/slog"

	"github.com/sirupsen/logrus"
)

type HandlerOptions struct {
	AddSource bool

	LevelMapper func(slog.Level) logrus.Level
}

type Handler struct {
	logger *logrus.Logger
	opts   HandlerOptions

	fields logrus.Fields

	groups []string
}

var _ slog.Handler = (*Handler)(nil)

func NewHandler(logger *logrus.Logger, opts *HandlerOptions) *Handler {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) toLogrusLevel(level slog.Level) logrus.Level {
	_ = "STUB: not implemented"
	return *new(logrus.Level)
}

func (h *Handler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *Handler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *Handler) clone() *Handler { _ = "STUB: not implemented"; return nil }

func appendAttr(fields logrus.Fields, groups []string, attr slog.Attr) {
	_ = "STUB: not implemented"
	return
}
