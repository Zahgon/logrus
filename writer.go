package logrus

import (
	"io"
)

func (logger *Logger) Writer() *io.PipeWriter { _ = "STUB: not implemented"; return nil }

func (logger *Logger) WriterLevel(level Level) *io.PipeWriter {
	_ = "STUB: not implemented"
	return nil
}

func (entry *Entry) Writer() *io.PipeWriter { _ = "STUB: not implemented"; return nil }

func (entry *Entry) WriterLevel(level Level) *io.PipeWriter { _ = "STUB: not implemented"; return nil }

func (entry *Entry) writerScanner(reader *io.PipeReader, printFunc func(args ...any)) {
	_ = "STUB: not implemented"
	return
}

func writerFinalizer(writer *io.PipeWriter) { _ = "STUB: not implemented"; return }
