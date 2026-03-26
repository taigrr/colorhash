package colorhash

import (
	"math"
	"testing"
)

func TestGenerateOKLCHPalette(t *testing.T) {
	palette := GenerateOKLCHPalette(8, 0.7, 0.15)

	if len(palette) != 8 {
		t.Fatalf("expected 8 colors, got %d", len(palette))
	}

	// All colors should have approximately the same lightness and chroma
	for i, c := range palette {
		oklch := c.ToOKLCH()
		if math.Abs(oklch.L-0.7) > 0.05 {
			t.Errorf("color %d: L = %f, expected ~0.7", i, oklch.L)
		}
		if math.Abs(oklch.C-0.15) > 0.03 {
			t.Errorf("color %d: C = %f, expected ~0.15", i, oklch.C)
		}
	}

	// Hues should be evenly spaced (45° apart for 8 colors)
	for i, c := range palette {
		oklch := c.ToOKLCH()
		expectedH := float64(i) * 45.0
		diff := math.Abs(oklch.H - expectedH)
		if diff > 180 {
			diff = 360 - diff
		}
		if diff > 10.0 {
			t.Errorf("color %d: H = %f, expected ~%f", i, oklch.H, expectedH)
		}
	}
}

func TestGenerateOKLCHPaletteEmpty(t *testing.T) {
	palette := GenerateOKLCHPalette(0, 0.7, 0.15)
	if len(palette) != 0 {
		t.Errorf("expected empty palette, got %d colors", len(palette))
	}
}

func TestGenerateOKLCHPaletteSingle(t *testing.T) {
	palette := GenerateOKLCHPalette(1, 0.6, 0.1)
	if len(palette) != 1 {
		t.Fatalf("expected 1 color, got %d", len(palette))
	}
	oklch := palette[0].ToOKLCH()
	if oklch.H > 5.0 && oklch.H < 355.0 {
		t.Errorf("single color H = %f, expected ~0", oklch.H)
	}
}

func TestGenerateOKLCHPaletteNegative(t *testing.T) {
	palette := GenerateOKLCHPalette(-5, 0.7, 0.15)
	if len(palette) != 0 {
		t.Errorf("expected empty palette for negative n, got %d colors", len(palette))
	}
}

func TestGenerateOKLCHPaletteLarge(t *testing.T) {
	palette := GenerateOKLCHPalette(360, 0.7, 0.15)
	if len(palette) != 360 {
		t.Fatalf("expected 360 colors, got %d", len(palette))
	}
	// Hue step should be ~1° for 360 colors
	c0 := palette[0].ToOKLCH()
	c1 := palette[1].ToOKLCH()
	step := c1.H - c0.H
	if math.Abs(step-1.0) > 1.0 {
		t.Errorf("expected ~1° hue step, got %f°", step)
	}
}

func TestGenerateOKLCHPaletteDistinct(t *testing.T) {
	palette := GenerateOKLCHPalette(6, 0.7, 0.15)
	// All colors should be distinct
	seen := make(map[int]bool)
	for _, c := range palette {
		if seen[int(c)] {
			t.Errorf("duplicate color: %s", c.ToHex())
		}
		seen[int(c)] = true
	}
}
