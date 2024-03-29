package colorhash

import (
	"fmt"
	"image/color"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red
)

var (
	Black   = ColorString("\033[1;30m%s\033[0m")
	Red     = ColorString("\033[1;31m%s\033[0m")
	Green   = ColorString("\033[1;32m%s\033[0m")
	Yellow  = ColorString("\033[1;33m%s\033[0m")
	Purple  = ColorString("\033[1;34m%s\033[0m")
	Magenta = ColorString("\033[1;35m%s\033[0m")
	Teal    = ColorString("\033[0;97m%s\033[0m")
	White   = ColorString("\033[1;37m%s\033[0m")
	//  Bold
	BBlack  = ColorString("\033[1;30m%s\033[0m")
	BRed    = ColorString("\033[1;31m%s\033[0m")
	BGreen  = ColorString("\033[1;32m%s\033[0m")
	BYellow = ColorString("\033[1;33m%s\033[0m")
	BBlue   = ColorString("\033[1;34m%s\033[0m")
	BPurple = ColorString("\033[1;35m%s\033[0m")
	BCyan   = ColorString("\033[1;36m%s\033[0m")
	BWhite  = ColorString("\033[1;37m%s\033[0m")

	//  Underline
	UBlack  = ColorString("\033[4;30m%s\033[0m")
	URed    = ColorString("\033[4;31m%s\033[0m")
	UGreen  = ColorString("\033[4;32m%s\033[0m")
	UYellow = ColorString("\033[4;33m%s\033[0m")
	UBlue   = ColorString("\033[4;34m%s\033[0m")
	UPurple = ColorString("\033[4;35m%s\033[0m")
	UCyan   = ColorString("\033[4;36m%s\033[0m")
	UWhite  = ColorString("\033[4;37m%s\033[0m")

	//  Background
	OnBlack  = ColorString("\033[40m%s\033[0m")
	OnRed    = ColorString("\033[41m%s\033[0m")
	OnGreen  = ColorString("\033[42m%s\033[0m")
	OnYellow = ColorString("\033[43m%s\033[0m")
	OnBlue   = ColorString("\033[44m%s\033[0m")
	OnPurple = ColorString("\033[45m%s\033[0m")
	OnCyan   = ColorString("\033[46m%s\033[0m")
	OnWhite  = ColorString("\033[47m%s\033[0m")

	//  High Intensty
	IBlack  = ColorString("\033[0;90m%s\033[0m")
	IRed    = ColorString("\033[0;91m%s\033[0m")
	IGreen  = ColorString("\033[0;92m%s\033[0m")
	IYellow = ColorString("\033[0;93m%s\033[0m")
	IBlue   = ColorString("\033[0;94m%s\033[0m")
	IPurple = ColorString("\033[0;95m%s\033[0m")
	ICyan   = ColorString("\033[0;96m%s\033[0m")
	IWhite  = ColorString("\033[0;97m%s\033[0m")

	//  Bold High Intensty
	BIBlack  = ColorString("\033[1;90m%s\033[0m")
	BIRed    = ColorString("\033[1;91m%s\033[0m")
	BIGreen  = ColorString("\033[1;92m%s\033[0m")
	BIYellow = ColorString("\033[1;93m%s\033[0m")
	BIBlue   = ColorString("\033[1;94m%s\033[0m")
	BIPurple = ColorString("\033[1;95m%s\033[0m")
	BICyan   = ColorString("\033[1;96m%s\033[0m")
	BIWhite  = ColorString("\033[1;97m%s\033[0m")

	//  High Intensty backgrounds
	OnIBlack  = ColorString("\033[0;100m%s\033[0m")
	OnIRed    = ColorString("\033[0;101m%s\033[0m")
	OnIGreen  = ColorString("\033[0;102m%s\033[0m")
	OnIYellow = ColorString("\033[0;103m%s\033[0m")
	OnIBlue   = ColorString("\033[0;104m%s\033[0m")
	OnIPurple = ColorString("\033[10;95m%s\033[0m")
	OnICyan   = ColorString("\033[0;106m%s\033[0m")
	OnIWhite  = ColorString("\033[0;107m%s\033[0m")
)

type ColorSet interface {
	ToPalette() color.Palette
	Get(int) color.Color
	Len() int
}

type StringerPalette []ColorStringer

func createStringerPalette(backgroundFillMode, disableSmartMode bool, c ...ColorSet) StringerPalette {
	palette := StringerPalette{}
	for _, colorSet := range c {
		for i := 0; i < colorSet.Len(); i++ {
			palette = append(palette, trueColorString(colorSet.Get(i), backgroundFillMode, disableSmartMode))
		}
	}
	return palette
}

// TBD not yet complete
func trueColorString(color color.Color, backgroundFillMode, disableSmartMode bool) ColorStringer {
	fgEsc, bgEsc := 48, 38
	sprint := func(args ...interface{}) string {
		r, g, b, _ := color.RGBA()
		if !disableSmartMode {
			return fmt.Sprintf("\033[;2;%d;%d;%d;m%s\033[0m\u001B[39m",
				r, g, b,
				fmt.Sprint(args...))
		}
		esc := fgEsc
		if backgroundFillMode {
			esc = bgEsc
		}

		return fmt.Sprintf("\033[%d;2;%d;%d;%d;m%s\033[0m\u001B[%dm",
			esc,
			r, g, b,
			fmt.Sprint(args...),
			esc+1)
	}
	return sprint
}

type ColorStringer func(...interface{}) string

func ColorString(colorString string) ColorStringer {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func GetBackgroundColor(c color.Color) color.Color {
	red, green, blue, _ := c.RGBA()
	if (float32(red)*0.299 + float32(green)*0.587 + float32(blue)*0.114) > 150.0 {
		return simplecolor.FromRGBA(0, 0, 0, 0)
	}
	return simplecolor.FromRGBA(255, 255, 255, 0)
}
