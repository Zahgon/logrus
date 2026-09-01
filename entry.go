package logrus

import (
	"bytes"
	"context"
	"runtime"
	"sync"
	"time"
)

var (
	logrusPackage string

	minimumCallerDepth = 1

	callerInitOnce sync.Once
)

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

var ErrorKey = "error"

//nolint:recvcheck // Entry methods intentionally use both pointer and value receivers.
type Entry struct {
	Logger *Logger

	Data Fields

	Time time.Time

	Level Level

	Caller *runtime.Frame

	Message string

	Buffer *bytes.Buffer

	Context context.Context

	err string
}

func NewEntry(logger *Logger) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) Dup() *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) dup() *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (entry *Entry) String() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (entry *Entry) WithError(err error) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) WithContext(ctx context.Context) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) WithField(key string, value any) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) WithFields(fields Fields) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) WithTime(t time.Time) *Entry { _ = "STUB: not implemented"; return nil }

func (entry *Entry) addField(key string, value any) { _ = "STUB: not implemented"; return }

func getPackageName(f string) string { _ = "STUB: not implemented"; return "" }

func getCaller() *runtime.Frame { _ = "STUB: not implemented"; return nil }

//go:fix inline
func (entry Entry) HasCaller() bool { _ = "STUB: not implemented"; return false }

func (entry *Entry) logArgs(level Level, panicAfter bool, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (entry *Entry) logf(level Level, panicAfter bool, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (entry *Entry) logln(level Level, panicAfter bool, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (entry *Entry) log(level Level, panicAfter bool, msg string) {
	_ = "STUB: not implemented"
	return
}

func (entry *Entry) getBufferPool() (pool BufferPool) {
	_ = "STUB: not implemented"
	return *new(BufferPool)
}

func (entry *Entry) fireHooks(hooks []Hook) { _ = "STUB: not implemented"; return }

func (entry *Entry) write() { _ = "STUB: not implemented"; return }

func (entry *Entry) Log(level Level, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Trace(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Debug(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Print(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Info(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warn(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warning(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Error(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Fatal(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Panic(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Logf(level Level, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (entry *Entry) Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Logln(level Level, args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Traceln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Debugln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Infoln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Println(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warnln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Warningln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Errorln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

func (entry *Entry) Panicln(args ...any) { _ = "STUB: not implemented"; return }

func sprint(args ...any) string { _ = "STUB: not implemented"; return "" }
