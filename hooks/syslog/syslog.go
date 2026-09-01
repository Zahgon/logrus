//go:build !windows && !nacl && !plan9

package syslog

import (
	"log/syslog"

	"github.com/sirupsen/logrus"
)

type SyslogHook struct {
	Writer        *syslog.Writer
	SyslogNetwork string
	SyslogRaddr   string
}

var _ logrus.Hook = (*SyslogHook)(nil)

func NewSyslogHook(network, raddr string, priority syslog.Priority, tag string) (*SyslogHook, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hook *SyslogHook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (hook *SyslogHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }
