package logger

import (
	"io"
	"os"
	"sync"
)

const DefaultPrefix = "DEFAULT"

var (
	loggerOnce map[string]*sync.Once
	Logs       map[string]*Logger
	Log        *Logger
)

func init() {
	loggerOnce = make(map[string]*sync.Once)
	Logs = make(map[string]*Logger)
	Logs[DefaultPrefix] = Default(DefaultPrefix)
	Log = Logs[DefaultPrefix]
}

type Logger struct {
	logFormatter      LoggingFormatter
	contentFormat     string
	Outs              []io.Writer
	Level             Level
	prefix            string
	forceEnableColor  bool
	forceDisableColor bool
}

func New(prefix string) *Logger {
	once := loggerOnce[prefix]
	if once == nil {
		once = &sync.Once{}
		loggerOnce[prefix] = once
	}
	once.Do(func() {
		Logs[prefix] = &Logger{prefix: prefix}
	})
	return Logs[prefix]
}

func Default(prefix string) *Logger {
	if logger, ok := Logs[prefix]; ok {
		return logger
	}
	logger := New(prefix)
	out := os.Stdout

	logger.Outs = append(logger.Outs, out)
	logger.Level = DebugLevel
	logger.logFormatter = &TextFormatter{}
	logger.SetContentFormat(DefaultContentFormat)
	return logger
}

func (logger *Logger) ForceColor() {
	logger.forceEnableColor = true
}

func (logger *Logger) DisColor() {
	logger.forceDisableColor = true
}

// SetFilePath
// 设置日志的输出路径
// 可以设置多个
// 默认会有一个标准输出
func (logger *Logger) SetFilePath(filePath string) {
	if filePath != "" {
		out, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
		logger.Outs = append(logger.Outs, out)
	}
}

func (logger *Logger) SetLogFormat(format LoggingFormatter) {
	logger.logFormatter = format
	logger.SetContentFormat(logger.contentFormat)
}

func (logger *Logger) SetContentFormat(format string) {
	logger.contentFormat = format
	str, contents := logger.ParseContentFormat(format)
	logger.logFormatter.SetContentFormatStr(str)
	logger.logFormatter.SetContents(contents)
}
