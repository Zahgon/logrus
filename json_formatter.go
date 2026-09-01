package logrus

import (
	"runtime"
)

type fieldKey string

type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string { _ = "STUB: not implemented"; return "" }

type JSONFormatter struct {
	TimestampFormat string

	DisableTimestamp bool

	DisableHTMLEscape bool

	DataKey string

	FieldMap FieldMap

	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	PrettyPrint bool
}

func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
