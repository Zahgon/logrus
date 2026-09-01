//go:build !appengine && !js && !windows && !nacl && !plan9 && !wasi && !wasip1 && !tinygo

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool { _ = "STUB: not implemented"; return false }
