package simple

import (
	"context"
	"github.com/Fullstack1014/logger"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := Default("test")
	logger.Warn("%s %s", "wendell", "25")
	logger.Debug("%s %s", "wendell", "25")
	logger.Info("%s %s", "wendell", "25")
	logger.Error("%s %s", "wendell", "25")
}

func TestLoggerLevel(t *testing.T) {
	logger := Default("test")
	logger.Logger.Level = InfoLevel
	logger.Debug("%s %s", "wendell", "25")
	logger.Info("%s %s", "wendell", "25")
	logger.Error("%s %s", "wendell", "25")
}

func TestValue(t *testing.T) {
	logger := Default("test")
	logger.Logger.Level = InfoLevel
	logger.Debug(ctx, "%s %s", "wendell", "25")
	logger.Info(ctx, "%s %s", "wendell", "25")
	logger.Error(ctx, "%s %s", "wendell", "25")
}

func TestLoggerToFile(t *testing.T) {
	todo := context.TODO()
	logger := Default("test")
	logger.Logger.SetFilePath("./logger.log")
	logger.Logger.ForceColor()
	logger.Debug(todo, "%s %s", "wendell", "25")
	logger.Info(todo, "%s %s", "wendell", "25")
	logger.Error(todo, "%s %s", "wendell", "25")
}

func TestLoggerJsonFormat(t *testing.T) {
	todo := context.TODO()
	log := Default("test")
	log.Logger.SetFilePath("./wits.log")
	log.Logger.SetLogFormat(logger.JsonFormat)
	log.Debug(todo, "%s %s", "wendell", "25")
	log.Info(todo, "%s %s", "wendell", "25")
	log.Error(todo, "%s %s", "wendell", "25")
}
