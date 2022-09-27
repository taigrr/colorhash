package colorhash

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	TotalHexColorspace = 16777216
)

// Extended ANSI colors (0-255).
var ansi = []int{
	0x000000, 0x800000, 0x008000, 0x808000, 0x000080, 0x800080, 0x008080,
	0xC0C0C0, 0x808080, 0xFF0000, 0x00FF00, 0xFFFF00, 0x0000FF, 0xFF00FF,
	0x00FFFF, 0xFFFFFF, 0x000000, 0x00005F, 0x000087, 0x0000AF, 0x0000D7,
	0x0000FF, 0x005F00, 0x005F5F, 0x005F87, 0x005FAF, 0x005FD7, 0x005FFF,
	0x008700, 0x00875F, 0x008787, 0x0087AF, 0x0087D7, 0x0087FF, 0x00AF00,
	0x00AF5F, 0x00AF87, 0x00AFAF, 0x00AFD7, 0x00AFFF, 0x00D700, 0x00D75F,
	0x00D787, 0x00D7AF, 0x00D7D7, 0x00D7FF, 0x00FF00, 0x00FF5F, 0x00FF87,
	0x00FFAF, 0x00FFD7, 0x00FFFF, 0x5F0000, 0x5F005F, 0x5F0087, 0x5F00AF,
	0x5F00D7, 0x5F00FF, 0x5F5F00, 0x5F5F5F, 0x5F5F87, 0x5F5FAF, 0x5F5FD7,
	0x5F5FFF, 0x5F8700, 0x5F875F, 0x5F8787, 0x5F87AF, 0x5F87D7, 0x5F87FF,
	0x5FAF00, 0x5FAF5F, 0x5FAF87, 0x5FAFAF, 0x5FAFD7, 0x5FAFFF, 0x5FD700,
	0x5FD75F, 0x5FD787, 0x5FD7AF, 0x5FD7D7, 0x5FD7FF, 0x5FFF00, 0x5FFF5F,
	0x5FFF87, 0x5FFFAF, 0x5FFFD7, 0x5FFFFF, 0x870000, 0x87005F, 0x870087,
	0x8700AF, 0x8700D7, 0x8700FF, 0x875F00, 0x875F5F, 0x875F87, 0x875FAF,
	0x875FD7, 0x875FFF, 0x878700, 0x87875F, 0x878787, 0x8787AF, 0x8787D7,
	0x8787FF, 0x87AF00, 0x87AF5F, 0x87AF87, 0x87AFAF, 0x87AFD7, 0x87AFFF,
	0x87D700, 0x87D75F, 0x87D787, 0x87D7AF, 0x87D7D7, 0x87D7FF, 0x87FF00,
	0x87FF5F, 0x87FF87, 0x87FFAF, 0x87FFD7, 0x87FFFF, 0xAF0000, 0xAF005F,
	0xAF0087, 0xAF00AF, 0xAF00D7, 0xAF00FF, 0xAF5F00, 0xAF5F5F, 0xaF5F87,
	0xAF5FAF, 0xAF5FD7, 0xAF5FFF, 0xAF8700, 0xAF875F, 0xaF8787, 0xAF87AF,
	0xAF87D7, 0xAF87FF, 0xAFAF00, 0xAFAF5F, 0xaFAF87, 0xAFAFAF, 0xAFAFD7,
	0xAFAFFF, 0xAFD700, 0xAFD75F, 0xaFD787, 0xAFD7AF, 0xAFD7D7, 0xAFD7FF,
	0xAFFF00, 0xAFFF5F, 0xaFFF87, 0xAFFFAF, 0xAFFFD7, 0xAFFFFF, 0xD70000,
	0xD7005F, 0xD70087, 0xD700AF, 0xD700D7, 0xD700FF, 0xD75F00, 0xD75F5F,
	0xD75F87, 0xD75FAF, 0xD75FD7, 0xD75FFF, 0xD78700, 0xD7875F, 0xD78787,
	0xD787AF, 0xD787D7, 0xD787FF, 0xD7AF00, 0xD7AF5F, 0xD7AF87, 0xD7AFAF,
	0xD7AFD7, 0xD7AFFF, 0xD7D700, 0xD7D75F, 0xD7D787, 0xD7D7AF, 0xD7D7D7,
	0xD7D7FF, 0xD7FF00, 0xD7FF5F, 0xD7FF87, 0xD7FFAF, 0xD7FFD7, 0xD7FFFF,
	0xFF0000, 0xFF005F, 0xFF0087, 0xFF00AF, 0xFF00D7, 0xFF00FF, 0xFF5F00,
	0xFF5F5F, 0xFF5F87, 0xFF5FAF, 0xFF5FD7, 0xFF5FFF, 0xFF8700, 0xFF875F,
	0xFF8787, 0xFF87AF, 0xFF87D7, 0xFF87FF, 0xFFAF00, 0xFFAF5F, 0xFFAF87,
	0xFFAFAF, 0xFFAFD7, 0xFFAFFF, 0xFFD700, 0xFFD75F, 0xFFD787, 0xFFD7AF,
	0xFFD7D7, 0xFFD7FF, 0xFFFF00, 0xFFFF5F, 0xFFFF87, 0xFFFFAF, 0xFFFFD7,
	0xFFFFFF, 0x080808, 0x121212, 0x1C1C1C, 0x262626, 0x303030, 0x3A3A3A,
	0x444444, 0x4E4E4E, 0x585858, 0x626262, 0x6C6C6C, 0x767676, 0x808080,
	0x8A8A8A, 0x949494, 0x9E9E9E, 0xA8A8A8, 0xB2B2B2, 0xBCBCBC, 0xC6C6C6,
	0xD0D0D0, 0xDADADA, 0xE4E4E4, 0xEEEEEE,
}

type Color struct {
	Hue   int
	Alpha int
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.GetRed()), uint32(c.GetGreen()), uint32(c.GetBlue()), uint32(c.Alpha)
}

func RGBA(r, g, b, a int) (c Color, err error) {
	if r > a {
		return c, errors.New("r value is greater than a value")
	}
	return Color{Hue: r<<16 + g<<8 + b}, nil
}

func RGB(r, g, b int) Color {
	return Color{Hue: r<<16 + g<<8 + b}
}

func (c Color) GetRed() int {
	return c.Hue >> 16 & 0xFF
}

func (c Color) GetGreen() int {
	return c.Hue >> 8 & 0xFF
}

func (c Color) GetBlue() int {
	return c.Hue & 0xFF
}

func CreateColor(x int) Color {
	return Color{Hue: x % TotalHexColorspace}
}

func (c Color) ToHex() string {
	return "#" + fmt.Sprintf("%06X", c.Hue)
}

func (c Color) ToShortHex() string {
	value := c.Hue >> 16 & 0xF
	value += c.Hue >> 8 & 0xF
	value += c.Hue & 0xF
	return "#" + fmt.Sprintf("%06X", value)
}

func FromHex(h string) (c Color, err error) {
	h = strings.TrimLeft(h, "#")
	h = strings.ToLower(h)
	if strings.HasPrefix(h, "0x") {
		h = strings.Replace(h, "0x", "", 1)
	}
	switch len(h) {
	case 3:
		var i int64
		i, err = strconv.ParseInt(string(h[0])+string(h[0]), 16, 64)
		if err != nil {
			return
		}
		fmt.Println(i)
		c.Hue += int(i << 16)
		i, err = strconv.ParseInt(string(h[1])+string(h[1]), 16, 64)
		if err != nil {
			return
		}
		c.Hue += int(i << 8)
		i, err = strconv.ParseInt(string(h[2])+string(h[2]), 16, 64)
		if err != nil {
			return
		}
		c.Hue += int(i)
	case 6:
		var i int64
		i, err = strconv.ParseInt(string(h[0:2]), 16, 64)
		if err != nil {
			return
		}
		fmt.Println(i)
		c.Hue += int(i << 16)
		i, err = strconv.ParseInt(string(h[2:4]), 16, 64)
		if err != nil {
			return
		}
		c.Hue += int(i << 8)
		i, err = strconv.ParseInt(string(h[4:6]), 16, 64)
		if err != nil {
			return
		}
		c.Hue += int(i)

	default:
		err = errors.New("invalid hex code length")
	}
	return
}

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

func ColorString(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
