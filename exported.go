package logrus

import (
	"context"
	"io"
	"time"
)

var std = New()

func StandardLogger() *Logger { _ = "STUB: not implemented"; return nil }

func SetOutput(out io.Writer) { _ = "STUB: not implemented"; return }

func SetFormatter(formatter Formatter) { _ = "STUB: not implemented"; return }

func SetReportCaller(include bool) { _ = "STUB: not implemented"; return }

func SetLevel(level Level) { _ = "STUB: not implemented"; return }

func GetLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

func IsLevelEnabled(level Level) bool { _ = "STUB: not implemented"; return false }

func AddHook(hook Hook) { _ = "STUB: not implemented"; return }

func WithError(err error) *Entry { _ = "STUB: not implemented"; return nil }

func WithContext(ctx context.Context) *Entry { _ = "STUB: not implemented"; return nil }

func WithField(key string, value any) *Entry { _ = "STUB: not implemented"; return nil }

func WithFields(fields Fields) *Entry { _ = "STUB: not implemented"; return nil }

func WithTime(t time.Time) *Entry { _ = "STUB: not implemented"; return nil }

func Trace(args ...any) { _ = "STUB: not implemented"; return }

func Debug(args ...any) { _ = "STUB: not implemented"; return }

func Print(args ...any) { _ = "STUB: not implemented"; return }

func Info(args ...any) { _ = "STUB: not implemented"; return }

func Warn(args ...any) { _ = "STUB: not implemented"; return }

func Warning(args ...any) { _ = "STUB: not implemented"; return }

func Error(args ...any) { _ = "STUB: not implemented"; return }

func Panic(args ...any) { _ = "STUB: not implemented"; return }

func Fatal(args ...any) { _ = "STUB: not implemented"; return }

func TraceFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func DebugFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func PrintFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func InfoFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func WarnFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func WarningFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func ErrorFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func PanicFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func FatalFn(fn LogFunction) { _ = "STUB: not implemented"; return }

func Tracef(format string, args ...any) { _ = "STUB: not implemented"; return }

func Debugf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Infof(format string, args ...any) { _ = "STUB: not implemented"; return }

func Warnf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Warningf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Panicf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Fatalf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Traceln(args ...any) { _ = "STUB: not implemented"; return }

func Debugln(args ...any) { _ = "STUB: not implemented"; return }

func Println(args ...any) { _ = "STUB: not implemented"; return }

func Infoln(args ...any) { _ = "STUB: not implemented"; return }

func Warnln(args ...any) { _ = "STUB: not implemented"; return }

func Warningln(args ...any) { _ = "STUB: not implemented"; return }

func Errorln(args ...any) { _ = "STUB: not implemented"; return }

func Panicln(args ...any) { _ = "STUB: not implemented"; return }

func Fatalln(args ...any) { _ = "STUB: not implemented"; return }
