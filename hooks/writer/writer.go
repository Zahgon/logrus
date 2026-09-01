package writer

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Hook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

var _ logrus.Hook = (*Hook)(nil)

func (hook *Hook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (hook *Hook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }
