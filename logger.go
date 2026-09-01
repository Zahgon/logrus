package logrus

import (
	"context"
	"io"
	"sync"
	"time"
)

type LogFunction func() []any

type Logger struct {
	Out io.Writer

	Hooks LevelHooks

	Formatter Formatter

	ReportCaller bool

	Level Level

	mu mutexWrap

	entryPool sync.Pool

	ExitFunc func(int)

	BufferPool BufferPool
}

type MutexWrap = mutexWrap

type mutexWrap struct {
	lock     sync.Mutex
	disabled bool
}

func (mw *mutexWrap) Lock() { _ = "STUB: not implemented"; return }

func (mw *mutexWrap) Unlock() { _ = "STUB: not implemented"; return }

func (mw *mutexWrap) Disable() { _ = "STUB: not implemented"; return }

func New() *Logger { _ = "STUB: not implemented"; return nil }

func (logger *Logger) newEntry() *Entry { _ = "STUB: not implemented"; return nil }

func (logger *Logger) releaseEntry(entry *Entry) { _ = "STUB: not implemented"; return }

func (logger *Logger) WithField(key string, value any) *Entry {
	_ = "STUB: not implemented"
	return nil
}

func (logger *Logger) WithFields(fields Fields) *Entry { _ = "STUB: not implemented"; return nil }

func (logger *Logger) WithError(err error) *Entry { _ = "STUB: not implemented"; return nil }

func (logger *Logger) WithContext(ctx context.Context) *Entry {
	_ = "STUB: not implemented"
	return nil
}

func (logger *Logger) WithTime(t time.Time) *Entry { _ = "STUB: not implemented"; return nil }

func (logger *Logger) Logf(level Level, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (logger *Logger) Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Log(level Level, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) LogFn(level Level, fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) Trace(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Debug(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Info(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Print(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warn(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warning(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Error(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Fatal(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Panic(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) TraceFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) DebugFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) InfoFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) PrintFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) WarnFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) WarningFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) ErrorFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) FatalFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) PanicFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func (logger *Logger) Logln(level Level, args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Traceln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Debugln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Infoln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Println(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warnln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Warningln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Errorln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Fatalln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Panicln(args ...any) { _ = "STUB: not implemented"; return }

func (logger *Logger) Exit(code int) { _ = "STUB: not implemented"; return }

func (logger *Logger) SetNoLock() { _ = "STUB: not implemented"; return }

func (logger *Logger) level() Level { _ = "STUB: not implemented"; return *new(Level) }

func (logger *Logger) SetLevel(level Level) { _ = "STUB: not implemented"; return }

func (logger *Logger) GetLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

func (logger *Logger) AddHook(hook Hook) { _ = "STUB: not implemented"; return }

func (logger *Logger) hooksForLevel(level Level) []Hook { _ = "STUB: not implemented"; return nil }

func (logger *Logger) IsLevelEnabled(level Level) bool { _ = "STUB: not implemented"; return false }

func (logger *Logger) SetFormatter(formatter Formatter) { _ = "STUB: not implemented"; return }

func (logger *Logger) SetOutput(output io.Writer) { _ = "STUB: not implemented"; return }

func (logger *Logger) SetReportCaller(reportCaller bool) { _ = "STUB: not implemented"; return }

func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks {
	_ = "STUB: not implemented"
	return *new(LevelHooks)
}

func (logger *Logger) SetBufferPool(pool BufferPool) { _ = "STUB: not implemented"; return }
