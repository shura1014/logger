package logger

import (
	"encoding/json"
	"fmt"
)

type JsonFormatter struct {
	BaseFormatter
}

func (f *JsonFormatter) Formatter(param *LoggingFormatterParam) string {
	params := make(map[string]any)
	for _, content := range f.contents {
		switch content.contentType {
		case fixedType:
			params[content.name] = content.content
		case timeType:
			params[TimeFiled] = NowFormat(content.name)
		case levelType:
			params[LevelFiled] = param.Level.Level()
		case msgType:
			params[MsgFiled] = param.Msg
		case lineType:
			// filter
		case paramType:
			params[content.content.(string)] = param.ctx.Value(content.content)
		default:
		}
	}
	marshal, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintln(string(marshal))
}
