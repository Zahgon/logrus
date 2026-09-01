package logrus

type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}

type LevelHooks map[Level][]Hook

func (hooks LevelHooks) Add(hook Hook) { _ = "STUB: not implemented"; return }

func (hooks LevelHooks) Fire(level Level, entry *Entry) error {
	_ = "STUB: not implemented"
	return nil
}
