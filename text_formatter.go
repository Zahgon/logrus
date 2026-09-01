package logrus

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var baseTimestamp = time.Now()

type TextFormatter struct {
	ForceColors bool

	DisableColors bool

	ForceQuote bool

	DisableQuote bool

	EnvironmentOverrideColors bool

	DisableTimestamp bool

	FullTimestamp bool

	TimestampFormat string

	DisableSorting bool

	SortingFunc func([]string)

	DisableLevelTruncation bool

	PadLevelText bool

	QuoteEmptyFields bool

	terminal bool

	FieldMap FieldMap

	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	terminalInitOnce sync.Once
}

func (f *TextFormatter) isTerminal(entry *Entry) bool { _ = "STUB: not implemented"; return false }

func (f *TextFormatter) isColored(isTerminal bool) bool { _ = "STUB: not implemented"; return false }

func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *TextFormatter) printPlain(b *bytes.Buffer, entry *Entry, keys []string, data Fields) {
	_ = "STUB: not implemented"
	return
}

func (f *TextFormatter) printColored(b *bytes.Buffer, entry *Entry, keys []string, data Fields) {
	_ = "STUB: not implemented"
	return
}

func (f *TextFormatter) appendKeyValue(b *bytes.Buffer, key string, value any) {
	_ = "STUB: not implemented"
	return
}

func (f *TextFormatter) appendValue(b *bytes.Buffer, value any) { _ = "STUB: not implemented"; return }

func (f *TextFormatter) appendString(b *bytes.Buffer, s string) { _ = "STUB: not implemented"; return }

func (f *TextFormatter) appendBytes(b *bytes.Buffer, bs []byte) { _ = "STUB: not implemented"; return }

func (f *TextFormatter) appendNumeric(b *bytes.Buffer, out []byte) {
	_ = "STUB: not implemented"
	return
}

func (f *TextFormatter) appendError(b *bytes.Buffer, v error) { _ = "STUB: not implemented"; return }

func (f *TextFormatter) appendStringer(b *bytes.Buffer, v fmt.Stringer) {
	_ = "STUB: not implemented"
	return
}

func (f *TextFormatter) recoverValue(b *bytes.Buffer, v any, method string) {
	_ = "STUB: not implemented"
	return
}

func needsQuoting(s string) bool { _ = "STUB: not implemented"; return false }

func needsQuotingBytes(bs []byte) bool { _ = "STUB: not implemented"; return false }

func isSafeByte(ch byte) bool { _ = "STUB: not implemented"; return false }
