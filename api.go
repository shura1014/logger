package logger

import (
	"context"
	"fmt"
	"io"
	"os"
)

func (logger *Logger) Info(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, InfoLevel, msg, GetFileNameAndLine(0), a...)
}

func (logger *Logger) InfoSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, InfoLevel, msg, GetFileNameAndLine(skip), a...)
}

func (logger *Logger) Debug(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, DebugLevel, msg, GetFileNameAndLine(0), a...)
}

func (logger *Logger) DebugSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, DebugLevel, msg, GetFileNameAndLine(skip), a...)
}

func (logger *Logger) Warn(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, WarnLevel, msg, GetFileNameAndLine(0), a...)
}

func (logger *Logger) WarnSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, WarnLevel, msg, GetFileNameAndLine(skip), a...)
}

func (logger *Logger) Error(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, ErrorLevel, msg, GetFileNameAndLine(0), a...)
}

func (logger *Logger) ErrorSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, ErrorLevel, msg, GetFileNameAndLine(skip), a...)
}

func (logger *Logger) TEXT(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, TEXT, msg, GetFileNameAndLine(0), a...)
}

func (logger *Logger) TEXTSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, TEXT, msg, GetFileNameAndLine(skip), a...)
}

// Fatal 将会退出程序
func (logger *Logger) Fatal(ctx context.Context, msg any, a ...any) {
	logger.DoPrint(ctx, ErrorLevel, msg, GetFileNameAndLine(0), a...)
	os.Exit(1)
}

func (logger *Logger) FatalSkip(ctx context.Context, msg any, skip int, a ...any) {
	logger.DoPrint(ctx, ErrorLevel, msg, GetFileNameAndLine(skip), a...)
	os.Exit(1)
}

// DoPrint
// level 日志级别
// msg日志内容
// a 日志内容的参数
// line 堆栈打印时需要跳过的行
// logger.Print(DebugLevel,"%s %s",0,"wendell",21)
func (logger *Logger) DoPrint(ctx context.Context, level Level, msg any, line string, a ...any) {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg.(string), a...)
	}
	if logger.logFormatter == nil {
		fmt.Println(msg)
		return
	}
	//级别不满足 不打印日志
	if logger.Level > level {
		return
	}
	// 可以支持多个输出，例如控制台输出加文本输出
	for _, out := range logger.Outs {
		color := false
		if !logger.forceDisableColor && (out == os.Stdout || logger.forceEnableColor) && logger.Level < TEXT {
			color = true
		} else {
			color = false
		}
		logger.Print(ctx, level, msg, line, color, out)

	}
}

// Print
// 日志打印
// level 日志的级别
// msg format格式
// a format格式的参数
func (logger *Logger) Print(ctx context.Context, level Level, msg any, line string, color bool, out io.Writer) {

	param := &LoggingFormatterParam{
		Level:  level,
		Msg:    msg,
		line:   line,
		prefix: logger.prefix,
		Color:  color,
		ctx:    ctx,
	}
	formatter := logger.logFormatter.Formatter(param)
	_, _ = fmt.Fprintf(out, "%v", formatter)
}
