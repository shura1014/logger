package logger

import "context"

type LoggingFormatter interface {
	Formatter(param *LoggingFormatterParam) string
	SetContentFormatStr(f string)
	SetContents(contents []*Content)
}

type LoggingFormatterParam struct {
	prefix string
	Msg    any
	line   string
	Color  bool
	Level  Level
	ctx    context.Context
}

var (
	DefaultFormat = &TextFormatter{}
	JsonFormat    = &JsonFormatter{}
)

type BaseFormatter struct {
	contentFormatStr string
	contents         []*Content
}

func (base *BaseFormatter) SetContentFormatStr(f string) {
	base.contentFormatStr = f
}

func (base *BaseFormatter) SetContents(contents []*Content) {
	base.contents = contents
}
