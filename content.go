package logger

import (
	"fmt"
	"log"
	"strings"
)

type ContentFunc func(msg ...any) any

type ContentType int

var (
	ContentFuncMap       map[string]ContentFunc
	DefaultContentFormat = `{[prefix::yellow]} {time:Y-M-D H:m:s::blue} {line} {level} {msg}`
	YmdhmsFiled          = "time:Y-M-D H:m:s"
	TimeFiled            = "time"
	LevelFiled           = "level"
	MsgFiled             = "msg"
	LineFiled            = "line"
	prefixFiled          = "prefix"
)

const (
	fixedType ContentType = iota
	timeType
	lineType
	levelType
	msgType
	paramType
)

func NothingFunc(msg ...any) any {
	return msg[0]
}

func init() {
	ContentFuncMap = make(map[string]ContentFunc)
	ContentFuncMap["yellow"] = YellowFunc
	ContentFuncMap["blue"] = BlueFunc
	ContentFuncMap["red"] = RedFunc
	ContentFuncMap["black"] = BlackFunc
	ContentFuncMap["green"] = GreenFunc
	ContentFuncMap["cyan"] = CyanFunc
	ContentFuncMap["grey"] = GreyFunc
	ContentFuncMap["time:Y-M-D H:m:s"] = TimeFunc("Y-M-D H:m:s")
}

// RegistryContentFunc 注册处理内容的函数
func RegistryContentFunc(name string, fun ContentFunc) {
	ContentFuncMap[name] = fun
}

// Content 根据 ContentFormat 解析
// {[prefix]} {time:Y-M-D H:m:s::blue} {line} {level} {msg}
type Content struct {
	name        string
	content     any
	contentType ContentType
	contentFunc ContentFunc
}

func NewContent() *Content {
	return &Content{}
}

// Wrap 包装ContentFunc
// 例如 time 需要先执行 TimeFunc -> BlueFunc
func (c *Content) Wrap(fun ContentFunc) {
	if c.contentFunc == nil {
		c.contentFunc = fun
	} else {
		contentFunc := c.contentFunc
		c.contentFunc = func(msg ...any) any {
			return fun(contentFunc(msg))
		}
	}
}

// ParseContentFormat 解析格式
func (logger *Logger) ParseContentFormat(s string) (f string, contents []*Content) {
	// 要求日志格式中每块都以一个空格隔开
	formats := strings.Split(s, "} {")
	var (
		formatStr strings.Builder
	)
	for _, format := range formats {

		// trim {}[]
		format = strings.ReplaceAll(format, "{", "")
		format = strings.ReplaceAll(format, "}", "")
		if format[0] == '[' {
			formatStr.WriteString("[%v] ")
			format = format[1 : len(format)-1]
		} else {
			formatStr.WriteString("%+v ")
		}
		content := NewContent()
		// 是否自定义函数
		index := strings.Index(format, "::")
		var (
			fun   ContentFunc
			value string
		)
		if index != -1 {
			value = format[0:index]
			fun1, ok := ContentFuncMap[format[index+2:]]
			// 如果自定义了函数，又没有注册函数，panic
			if !ok {
				log.Panicf("No function %s was found available", format[index+2:])
			}
			fun = fun1
		} else {
			value = format
		}

		switch value {
		case prefixFiled:
			// 固定不变的值
			content.contentType = fixedType
			content.name = prefixFiled
			value = logger.prefix
		case YmdhmsFiled:
			content.contentType = timeType
			value = Parse(value[5:])
			content.name = value
			content.Wrap(TimeFunc(value))
		case LevelFiled:
			content.contentType = levelType
			content.Wrap(LevelFunc)
			value = LevelFiled
		case MsgFiled:
			content.contentType = msgType
			content.Wrap(MsgFunc)
			value = MsgFiled
		case LineFiled:
			content.contentType = lineType
		default:
			if strings.HasPrefix(value, "time:") {
				content.contentType = timeType
				content.Wrap(TimeFunc(value[5:]))
				value = TimeFiled
			}

			if strings.HasPrefix(value, "param.") {
				content.contentType = paramType
				value = value[6:]
			}
		}

		content.content = value

		if fun != nil {
			content.Wrap(fun)
		}

		if content.contentFunc == nil {
			content.contentFunc = NothingFunc
		}

		//if content.contentType == fixedType {
		//	content.content = content.contentFunc(content.content)
		//}
		contents = append(contents, content)
	}
	formatStr.WriteByte('\n')
	return formatStr.String(), contents
}

func YellowFunc(msg ...any) any {
	return Yellow(msg[0])
}

func BlueFunc(msg ...any) any {
	return Blue(msg[0])
}

func BlackFunc(msg ...any) any {
	return Black(msg[0])
}

func RedFunc(msg ...any) any {
	return Red(msg[0])
}

func GreenFunc(msg ...any) any {
	return Green(msg[0])
}

func CyanFunc(msg ...any) any {
	return Cyan(msg[0])
}

func GreyFunc(msg ...any) any {
	return Grey(msg[0])
}

func TimeFunc(f string) ContentFunc {
	parse := Parse(f)
	return func(msg ...any) any {
		return NowFormat(parse)
	}
}

func LevelFunc(args ...any) any {
	level := args[0].(Level)
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

func MsgFunc(args ...any) any {
	level := args[0].(Level)
	msg := args[1]
	switch level {
	case WarnLevel:
		return Yellow(msg)
	case DebugLevel:
		return msg
	case InfoLevel:
		return msg
	case ErrorLevel:
		return Red(fmt.Sprintf(" Error Cause by: \n\t %+v", Red(msg)))
	default:
		return Cyan(msg)
	}
}

func IsArray(a []string, s string) bool {
	for _, v := range a {
		if s == v {
			return true
		}
	}
	return false
}
