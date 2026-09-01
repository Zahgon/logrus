package logrus

import (
	"bytes"
	"sync"
)

var bufferPool BufferPool = &defaultPool{
	pool: &sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	},
}

type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}

type defaultPool struct {
	pool *sync.Pool
}

func (p *defaultPool) Put(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (p *defaultPool) Get() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func SetBufferPool(bp BufferPool) { _ = "STUB: not implemented"; return }
