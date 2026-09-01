package logrus

import (
	"log"
)

type Fields map[string]any

//nolint:recvcheck // the methods of "Entry" use pointer receiver and non-pointer receiver.
type Level uint32

func (level Level) String() string { _ = "STUB: not implemented"; return "" }

func ParseLevel(lvl string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

func parseLevel(b []byte) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

func (level *Level) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (level Level) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

const (
	PanicLevel Level = iota

	FatalLevel

	ErrorLevel

	WarnLevel

	InfoLevel

	DebugLevel

	TraceLevel
)

var (
	_ StdLogger = (*log.Logger)(nil)
	_ StdLogger = (*Entry)(nil)
	_ StdLogger = (*Logger)(nil)

	_ FieldLogger = (*Logger)(nil)
	_ FieldLogger = (*Entry)(nil)
	_ FieldLogger = Ext1FieldLogger(nil)

	_ DebugLogger = (*Logger)(nil)
	_ InfoLogger  = (*Logger)(nil)
	_ WarnLogger  = (*Logger)(nil)
	_ ErrorLogger = (*Logger)(nil)
	_ TraceLogger = (*Logger)(nil)

	_ DebugLogger = (*Entry)(nil)
	_ InfoLogger  = (*Entry)(nil)
	_ WarnLogger  = (*Entry)(nil)
	_ ErrorLogger = (*Entry)(nil)
	_ TraceLogger = (*Entry)(nil)

	_ Ext1FieldLogger = (*Logger)(nil)
	_ Ext1FieldLogger = (*Entry)(nil)
)

type StdLogger interface {
	Print(args ...any)
	Printf(format string, args ...any)
	Println(args ...any)

	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Fatalln(args ...any)

	Panic(args ...any)
	Panicf(format string, args ...any)
	Panicln(args ...any)
}

type FieldLogger interface {
	WithField(key string, value any) *Entry
	WithFields(fields Fields) *Entry
	WithError(err error) *Entry

	StdLogger
	DebugLogger
	InfoLogger
	WarnLogger
	ErrorLogger

	Warning(args ...any)
	Warningf(format string, args ...any)
	Warningln(args ...any)
}

type DebugLogger interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Debugln(args ...any)
}

type InfoLogger interface {
	Info(args ...any)
	Infof(format string, args ...any)
	Infoln(args ...any)
}

type WarnLogger interface {
	Warn(args ...any)
	Warnf(format string, args ...any)
	Warnln(args ...any)
}

type ErrorLogger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
	Errorln(args ...any)
}

type TraceLogger interface {
	Trace(args ...any)
	Tracef(format string, args ...any)
	Traceln(args ...any)
}

type Ext1FieldLogger interface {
	FieldLogger
	TraceLogger
}
