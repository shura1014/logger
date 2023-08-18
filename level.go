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
