package logrus

import "time"

const (
	defaultTimestampFormat = time.RFC3339

	defaultFields = 3
)

const (
	FieldKeyMsg         = "msg"
	FieldKeyLevel       = "level"
	FieldKeyTime        = "time"
	FieldKeyLogrusError = "logrus_error"
	FieldKeyFunc        = "func"
	FieldKeyFile        = "file"
)

type Formatter interface {
	Format(*Entry) ([]byte, error)
}

func prefixFieldClashes(data Fields, fieldMap FieldMap, reportCaller bool) {
	_ = "STUB: not implemented"
	return
}
