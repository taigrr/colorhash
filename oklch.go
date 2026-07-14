// Package colorhash maps arbitrary strings and byte streams to
// deterministic colors from a given palette.
package colorhash

import (
	"math"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

// GenerateOKLCHPalette generates n evenly-spaced colors in the OKLCH color space
// at the given lightness and chroma. This produces a perceptually uniform palette
// where all colors appear equally bright and saturated.
func GenerateOKLCHPalette(n int, l, c float64) simplecolor.SimplePalette {
	return GenerateOKLCHPaletteWithHueOffset(n, l, c, 0)
}

// GenerateOKLCHPaletteWithHueOffset generates n evenly-spaced colors in the
// OKLCH color space, starting at startHue degrees.
func GenerateOKLCHPaletteWithHueOffset(n int, l, c, startHue float64) simplecolor.SimplePalette {
	if n <= 0 {
		return simplecolor.SimplePalette{}
	}
	palette := make(simplecolor.SimplePalette, n)
	step := 360.0 / float64(n)
	for i := 0; i < n; i++ {
		h := normalizeHue(startHue + float64(i)*step)
		palette[i] = simplecolor.FromOKLCH(l, c, h)
	}
	return palette
}

func normalizeHue(h float64) float64 {
	h = math.Mod(h, 360)
	if h < 0 {
		h += 360
	}
	return h
}
