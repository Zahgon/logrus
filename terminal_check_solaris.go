//go:build solaris && !tinygo

package logrus

func isTerminal(fd int) bool { _ = "STUB: not implemented"; return false }
