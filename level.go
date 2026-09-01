package logrus

import (
	"sync"
)

const (
	ansiReset    = "\x1b[0m"
	ansiRed      = "\x1b[31m"
	ansiYellow   = "\x1b[33m"
	ansiCyan     = "\x1b[36m"
	ansiDimCyan  = "\x1b[2;36m"
	ansiDimWhite = "\x1b[2;37m"
)

type lvlPrefix struct {
	full      string
	truncated string
	padded    string
}

func colorize(level Level, s string) string { _ = "STUB: not implemented"; return "" }

func formatLevel(level Level, disableTrunc, pad bool, maxLen int) string {
	_ = "STUB: not implemented"
	return ""
}

var levelPrefixOnce = sync.OnceValues(func() (map[Level]lvlPrefix, lvlPrefix) {
	var maxLevel Level
	maxLen := 0
	for _, lvl := range AllLevels {
		if lvl > maxLevel {
			maxLevel = lvl
		}
		if l := len(lvl.String()); l > maxLen {
			maxLen = l
		}
	}

	prefix := make(map[Level]lvlPrefix, len(AllLevels))
	for _, lvl := range AllLevels {
		prefix[lvl] = lvlPrefix{
			full:      formatLevel(lvl, true, false, maxLen),
			truncated: formatLevel(lvl, false, false, maxLen),
			padded:    formatLevel(lvl, true, true, maxLen),
		}
	}

	unknownLevel := maxLevel + 1
	unknown := lvlPrefix{
		full:      formatLevel(unknownLevel, true, false, maxLen),
		truncated: formatLevel(unknownLevel, false, false, maxLen),
		padded:    formatLevel(unknownLevel, true, true, maxLen),
	}

	return prefix, unknown
})

func levelPrefix(level Level, disableTrunc, pad bool) string { _ = "STUB: not implemented"; return "" }
