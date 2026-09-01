//go:build (darwin || dragonfly || freebsd || netbsd || openbsd || hurd) && !tinygo

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

func isTerminal(fd int) bool { _ = "STUB: not implemented"; return false }
