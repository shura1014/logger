package logger

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	TEXT
)

func (level Level) Level() string {
	switch level {
	case WarnLevel:
		return "WARN"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case ErrorLevel:
		return "ERROR"
	default:
		return ""
	}
}

func LevelColor(level Level) string {
	switch level {
	case WarnLevel:
		return Yellow(level.Level())
	case DebugLevel:
		return Grey(level.Level())
	case InfoLevel:
		return Green(level.Level())
	case ErrorLevel:
		return Red(level.Level())
	default:
		return reset(level.Level())
	}
}

func MsgColor(level Level, msg any) any {
	switch level {
	case WarnLevel:
		return Yellow(msg)
	case DebugLevel:
		return msg
	case InfoLevel:
		return msg
	case ErrorLevel:
		return Red(msg)
	default:
		return Cyan(msg)
	}
}
