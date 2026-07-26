package colorhash_test

import (
	"fmt"
	"image/color"

	"github.com/taigrr/colorhash"
	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

type examplePalette []color.Color

func (p examplePalette) ToPalette() color.Palette { return color.Palette(p) }
func (p examplePalette) Get(i int) color.Color    { return p[i] }
func (p examplePalette) Len() int                 { return len(p) }

func ExampleHashString() {
	fmt.Println(colorhash.HashString("hello colorhash"))
	// Output: 893132354324239557
}

func ExampleStringToColor() {
	palette := examplePalette{
		simplecolor.FromRGBA(255, 0, 0, 255),
		simplecolor.FromRGBA(0, 255, 0, 255),
		simplecolor.FromRGBA(0, 0, 255, 255),
	}

	c := colorhash.StringToColor(palette, "alice")
	fmt.Println(c == nil)
	// Output: false
}

func ExampleGenerateOKLCHPalette() {
	palette := colorhash.GenerateOKLCHPalette(8, 0.7, 0.15)
	fmt.Println(palette.Len())
	// Output: 8
}
