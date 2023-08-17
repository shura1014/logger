package simple

import (
	"context"
	"github.com/Fullstack1014/logger"
	"os"
)

const (
	DebugLevel = logger.DebugLevel

	InfoLevel  = logger.InfoLevel
	WarnLevel  = logger.WarnLevel
	ErrorLevel = logger.ErrorLevel
	TEXT       = logger.TEXT
)

var ctx = context.TODO()

// Log 很多时候不需要使用到ctx，加了ctx后代码看起来比较凌乱，提供一个不需要传入ctx的log
type Log struct {
	Logger *logger.Logger
}

func New(prefix string) *Log {
	return &Log{Logger: logger.New(prefix)}
}

func Default(prefix string) *Log {
	return &Log{Logger: logger.Default(prefix)}
}

func (log *Log) Info(msg any, a ...any) {
	log.Logger.DoPrint(ctx, InfoLevel, msg, logger.GetFileNameAndLine(0), a...)
}

func (log *Log) InfoSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, InfoLevel, msg, logger.GetFileNameAndLine(skip), a...)
}

func (log *Log) Debug(msg any, a ...any) {
	log.Logger.DoPrint(ctx, DebugLevel, msg, logger.GetFileNameAndLine(0), a...)
}

func (log *Log) DebugSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, DebugLevel, msg, logger.GetFileNameAndLine(skip), a...)
}

func (log *Log) Warn(msg any, a ...any) {
	log.Logger.DoPrint(ctx, WarnLevel, msg, logger.GetFileNameAndLine(0), a...)
}

func (log *Log) WarnSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, WarnLevel, msg, logger.GetFileNameAndLine(skip), a...)
}

func (log *Log) Error(msg any, a ...any) {
	log.Logger.DoPrint(ctx, ErrorLevel, msg, logger.GetFileNameAndLine(0), a...)
}

func (log *Log) ErrorSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, ErrorLevel, msg, logger.GetFileNameAndLine(skip), a...)
}

func (log *Log) TEXT(msg any, a ...any) {
	log.Logger.DoPrint(ctx, TEXT, msg, logger.GetFileNameAndLine(0), a...)
}

func (log *Log) TEXTSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, TEXT, msg, logger.GetFileNameAndLine(skip), a...)
}

// Fatal 将会退出程序
func (log *Log) Fatal(msg any, a ...any) {
	log.Logger.DoPrint(ctx, ErrorLevel, msg, logger.GetFileNameAndLine(0), a...)
	os.Exit(1)
}

func (log *Log) FatalSkip(msg any, skip int, a ...any) {
	log.Logger.DoPrint(ctx, ErrorLevel, msg, logger.GetFileNameAndLine(skip), a...)
	os.Exit(1)
}
