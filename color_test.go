package logger

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(Green("Green"))
	fmt.Println(Green("Green", Reset))
	fmt.Println(Green("Green", Italic))
	fmt.Println(Black("Black"))
	fmt.Println(Red("Red"))
	fmt.Println(Blue("Blue"))
	fmt.Println(Magenta("Magenta"))
	fmt.Println(Yellow("Yellow"))
	fmt.Println(Cyan("Cyan"))
	fmt.Println(White("White"))
	fmt.Println(Grey("Grey"))
	fmt.Println(blackBg("blackBg"))
	fmt.Println(redBg("redBg"))
	fmt.Println(greenBg("greenBg"))
	fmt.Println(yellowBg("yellowBg"))
	fmt.Println(blueBg("blueBg"))
	fmt.Println(magentaBg("magentaBg"))
	fmt.Println(cyanBg("cyanBg"))
	fmt.Println(whiteBg("whiteBg"))
	fmt.Println(bold("bold"))
	fmt.Println(dim("dim"))
	fmt.Println(italic("italic"))
	fmt.Println(underline("underline"))
	fmt.Println(hidden("hidden"))
	fmt.Println(reset("reset"))
	fmt.Println(inverse("inverse"))
	fmt.Println(strikeout("strikeout"))
}
