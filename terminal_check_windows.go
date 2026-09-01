//go:build windows && !appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool { _ = "STUB: not implemented"; return false }
