package test

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Hook struct {
	Entries []logrus.Entry
	mu      sync.RWMutex
}

var _ logrus.Hook = (*Hook)(nil)

func NewGlobal() *Hook { _ = "STUB: not implemented"; return nil }

func NewLocal(logger *logrus.Logger) *Hook { _ = "STUB: not implemented"; return nil }

func NewNullLogger() (*logrus.Logger, *Hook) { _ = "STUB: not implemented"; return nil, nil }

func (t *Hook) Fire(e *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (t *Hook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (t *Hook) LastEntry() *logrus.Entry { _ = "STUB: not implemented"; return nil }

func (t *Hook) AllEntries() []*logrus.Entry { _ = "STUB: not implemented"; return nil }

func (t *Hook) Reset() { _ = "STUB: not implemented"; return }
