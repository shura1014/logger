package logger

import (
	"bytes"
	"fmt"
)

type (
	inner func(interface{}, ...string) string
)

// Color styles
const (
	// Blk 黑色
	/****************************颜色 color style****************************/
	Blk = "30"
	// Rd 红色
	Rd = "31"
	// Grn 绿色
	Grn = "32"
	// Yel 黄色
	Yel = "33"
	// Blu 蓝色
	Blu = "34"
	// Mgn 品红色的
	Mgn = "35"
	// Cyn 青色
	Cyn = "36"
	// Wht 白色
	Wht = "37"
	// Gry 灰色
	Gry = "90"

	// BlkBg Black background style
	/****************************颜色背景 background style****************************/
	BlkBg = "40"
	RdBg  = "41"
	GrnBg = "42"
	YelBg = "43"
	BluBg = "44"
	MgnBg = "45"
	CynBg = "46"
	WhtBg = "47"

	// Reset
	/****************************样式 Here is the style****************************/
	Reset = "0"
	// Bold  粗体
	Bold = "1"
	// Dim 淡
	Dim = "2"
	// Italic 斜体
	Italic = "3"
	// Underline 下划线
	Underline = "4"
	// Inverse 背景反色
	Inverse = "7"
	// Hidden 隐藏
	Hidden = "8"
	// Strikeout 加删除线
	Strikeout = "9"
)

var (
	Black   = outer(Blk)
	Red     = outer(Rd)
	Green   = outer(Grn)
	Yellow  = outer(Yel)
	Blue    = outer(Blu)
	Magenta = outer(Mgn)
	Cyan    = outer(Cyn)
	White   = outer(Wht)
	Grey    = outer(Gry)

	blackBg   = outer(BlkBg)
	redBg     = outer(RdBg)
	greenBg   = outer(GrnBg)
	yellowBg  = outer(YelBg)
	blueBg    = outer(BluBg)
	magentaBg = outer(MgnBg)
	cyanBg    = outer(CynBg)
	whiteBg   = outer(WhtBg)

	reset     = outer(Reset)
	bold      = outer(Bold)
	dim       = outer(Dim)
	italic    = outer(Italic)
	underline = outer(Underline)
	inverse   = outer(Inverse)
	hidden    = outer(Hidden)
	strikeout = outer(Strikeout)
)

func outer(first string) inner {
	return func(msg interface{}, styles ...string) string {
		b := new(bytes.Buffer)
		b.WriteString("\x1b[")
		b.WriteString(first)
		for _, s := range styles {
			b.WriteString(";")
			b.WriteString(s)
		}
		b.WriteString("m")
		return fmt.Sprintf("%s%v\x1b[0m", b.String(), msg)
	}
}
