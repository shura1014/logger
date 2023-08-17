package logger

import (
	"fmt"
)

type TextFormatter struct {
	BaseFormatter
}

func (f *TextFormatter) Formatter(param *LoggingFormatterParam) string {
	if param.Color {
		args := make([]any, len(f.contents))
		for index, content := range f.contents {
			switch content.contentType {
			case fixedType:
				args[index] = content.contentFunc(content.content)
			case timeType:
				args[index] = content.contentFunc()
			case levelType:
				args[index] = content.contentFunc(param.Level)
			case msgType:
				args[index] = content.contentFunc(param.Level, param.Msg)
			case lineType:
				args[index] = param.line
			case paramType:
				value := param.ctx.Value(content.content)
				args[index] = content.contentFunc(value)

			default:
			}
		}
		return fmt.Sprintf(f.contentFormatStr, args...)
	}

	args := make([]any, len(f.contents))
	for index, content := range f.contents {
		switch content.contentType {
		case fixedType:
			args[index] = content.content
		case timeType:
			args[index] = NowFormat(content.name)
		case levelType:
			args[index] = param.Level.Level()
		case msgType:
			args[index] = param.Msg
		case lineType:
			args[index] = param.line
		case paramType:
			args[index] = param.ctx.Value(content.content)
		default:
		}
	}
	return fmt.Sprintf(f.contentFormatStr, args...)
}
