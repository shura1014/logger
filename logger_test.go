package logger

import (
	"context"
	"testing"
)

func TestLogger(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	//logger.DisColor()
	logger.Warn(todo, "%s %s", "wendell", "25")
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}

func TestLoggerLevel(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.Level = InfoLevel
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}

func TestValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "STAFF_ID", "TEST01")
	logger := Default("test")
	logger.SetContentFormat(`{[prefix::yellow]} {time:Y-M-D H:m:s::blue} {[param.STAFF_ID::red]} {line} {level} {msg}`)
	logger.Level = InfoLevel
	logger.Debug(ctx, "%s %s", "wendell", "25")
	logger.Info(ctx, "%s %s", "wendell", "25")
	logger.Error(ctx, "%s %s", "wendell", "25")
}

func TestLoggerToFile(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.SetFilePath("./logger.log")
	logger.ForceColor()
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}

func TestLoggerJsonFormat(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	//logger.SetFilePath("./logger.log")
	logger.SetLogFormat(JsonFormat)
	logger.DisColor()
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}
