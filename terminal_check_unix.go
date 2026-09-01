//go:build (linux || aix || zos) && !tinygo

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS

func isTerminal(fd int) bool { _ = "STUB: not implemented"; return false }
