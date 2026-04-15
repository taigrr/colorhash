package colorhash

import (
	"bytes"
	"fmt"
	"image/color"
	"testing"

	"github.com/taigrr/simplecolorpalettes/simplecolor"
)

// testPalette wraps a slice of colors to satisfy ColorSet.
type testPalette []color.Color

func (p testPalette) ToPalette() color.Palette { return color.Palette(p) }
func (p testPalette) Get(i int) color.Color    { return p[i] }
func (p testPalette) Len() int                 { return len(p) }

func newTestPalette() testPalette {
	return testPalette{
		simplecolor.FromRGBA(255, 0, 0, 255),
		simplecolor.FromRGBA(0, 255, 0, 255),
		simplecolor.FromRGBA(0, 0, 255, 255),
		simplecolor.FromRGBA(255, 255, 0, 255),
		simplecolor.FromRGBA(255, 0, 255, 255),
		simplecolor.FromRGBA(0, 255, 255, 255),
		simplecolor.FromRGBA(128, 128, 128, 255),
		simplecolor.FromRGBA(255, 128, 0, 255),
	}
}

func TestHashString(t *testing.T) {
	testStrings := []struct {
		String string
		Value  int
		ID     string
	}{
		{String: "", Value: 5472609002491880228, ID: "Empty string"},
		{String: "123", Value: 6449148174219763898, ID: "123"},
		{String: "it's as easy as", Value: 5908178111834329190, ID: "easy"},
		{String: "hello colorhash", Value: 893132354324239557, ID: "hello"},
	}
	for _, tc := range testStrings {
		t.Run(tc.ID, func(t *testing.T) {
			hash := HashString(tc.String)
			if hash != tc.Value {
				t.Errorf("%s :: Hash resulted in value %d, but expected value is %d", tc.ID, hash, tc.Value)
			}
		})
	}
}

func TestHashStringDeterministic(t *testing.T) {
	input := "deterministic-test"
	h1 := HashString(input)
	h2 := HashString(input)
	if h1 != h2 {
		t.Errorf("HashString is not deterministic: %d != %d", h1, h2)
	}
}

func TestHashStringPositive(t *testing.T) {
	inputs := []string{"", "a", "test", "negative?", "🎨"}
	for _, s := range inputs {
		h := HashString(s)
		if h < 0 {
			t.Errorf("HashString(%q) returned negative value: %d", s, h)
		}
	}
}

func TestHashBytes(t *testing.T) {
	input := []byte("hello colorhash")
	h := HashBytes(bytes.NewReader(input))
	expected := HashString("hello colorhash")
	if h != expected {
		t.Errorf("HashBytes and HashString diverged for same input: %d != %d", h, expected)
	}
}

func TestHashBytesDeterministic(t *testing.T) {
	input := []byte("deterministic")
	h1 := HashBytes(bytes.NewReader(input))
	h2 := HashBytes(bytes.NewReader(input))
	if h1 != h2 {
		t.Errorf("HashBytes is not deterministic: %d != %d", h1, h2)
	}
}

func TestStringToColor(t *testing.T) {
	palette := newTestPalette()
	c := StringToColor(palette, "test")
	if c == nil {
		t.Fatal("StringToColor returned nil")
	}
}

func TestBytesToColor(t *testing.T) {
	palette := newTestPalette()
	c := BytesToColor(palette, bytes.NewReader([]byte("test")))
	if c == nil {
		t.Fatal("BytesToColor returned nil")
	}
}

func TestStringToColorDeterministic(t *testing.T) {
	palette := newTestPalette()
	c1 := StringToColor(palette, "consistent")
	c2 := StringToColor(palette, "consistent")
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
		t.Error("StringToColor is not deterministic")
	}
}

func TestColorString(t *testing.T) {
	result := Red("error")
	if result == "" {
		t.Fatal("ColorString returned empty string")
	}
	if result == "error" {
		t.Fatal("ColorString did not wrap with escape codes")
	}
}

func TestGetBackgroundColor(t *testing.T) {
	// White text should get black background
	white := simplecolor.FromRGBA(255, 255, 255, 0)
	bg := GetBackgroundColor(white)
	r, g, b, _ := bg.RGBA()
	if r != 0 || g != 0 || b != 0 {
		t.Errorf("expected black background for white, got (%d,%d,%d)", r, g, b)
	}

	// Black text should get white background
	black := simplecolor.FromRGBA(0, 0, 0, 0)
	bg2 := GetBackgroundColor(black)
	r2, g2, b2, _ := bg2.RGBA()
	if r2 == 0 && g2 == 0 && b2 == 0 {
		t.Error("expected white background for black, got black")
	}
}

func TestStringerPalette(t *testing.T) {
	palette := newTestPalette()
	sp := createStringerPalette(false, false, palette)
	if len(sp) == 0 {
		t.Fatal("createStringerPalette returned empty palette")
	}
	result := sp.GetString("test")
	if result == "" {
		t.Fatal("GetString returned empty string")
	}
}

func TestCreateStringerPaletteBackgroundFill(t *testing.T) {
	palette := newTestPalette()
	sp := createStringerPalette(true, false, palette)
	if len(sp) != len(palette) {
		t.Fatalf("expected %d entries, got %d", len(palette), len(sp))
	}
	result := sp.GetString("test-bg")
	if result == "" {
		t.Fatal("GetString with background fill returned empty string")
	}
	if result == "test-bg" {
		t.Fatal("GetString with background fill did not wrap with escape codes")
	}
}

func TestCreateStringerPaletteDisableSmart(t *testing.T) {
	palette := newTestPalette()
	sp := createStringerPalette(false, true, palette)
	if len(sp) != len(palette) {
		t.Fatalf("expected %d entries, got %d", len(palette), len(sp))
	}
	result := sp.GetString("test-nosmart")
	if result == "" {
		t.Fatal("GetString with smart mode disabled returned empty string")
	}
}

func TestCreateStringerPaletteMultipleSets(t *testing.T) {
	p1 := newTestPalette()
	p2 := testPalette{
		simplecolor.FromRGBA(100, 100, 100, 255),
		simplecolor.FromRGBA(200, 200, 200, 255),
	}
	sp := createStringerPalette(false, false, p1, p2)
	if len(sp) != len(p1)+len(p2) {
		t.Fatalf("expected %d entries, got %d", len(p1)+len(p2), len(sp))
	}
}

func TestGetBackgroundColorMidTone(t *testing.T) {
	// A mid-tone color to exercise the luminance threshold
	mid := simplecolor.FromRGBA(128, 128, 128, 255)
	bg := GetBackgroundColor(mid)
	// Should return a valid color (either black or white)
	r, g, b, _ := bg.RGBA()
	isBlack := r == 0 && g == 0 && b == 0
	isWhite := r == 255 && g == 255 && b == 255
	if !isBlack && !isWhite {
		t.Errorf("expected black or white background, got (%d,%d,%d)", r, g, b)
	}
}

func TestColorStringVariants(t *testing.T) {
	variants := []struct {
		name string
		fn   ColorStringer
	}{
		{"Green", Green},
		{"Yellow", Yellow},
		{"Purple", Purple},
		{"Magenta", Magenta},
		{"BBlue", BBlue},
		{"UCyan", UCyan},
		{"OnRed", OnRed},
		{"IBlue", IBlue},
		{"BICyan", BICyan},
		{"OnIGreen", OnIGreen},
	}
	for _, v := range variants {
		t.Run(v.name, func(t *testing.T) {
			result := v.fn("test")
			if result == "" {
				t.Fatalf("%s returned empty string", v.name)
			}
			if result == "test" {
				t.Fatalf("%s did not apply escape codes", v.name)
			}
		})
	}
}

func TestHashBytesEmpty(t *testing.T) {
	h := HashBytes(bytes.NewReader([]byte{}))
	if h < 0 {
		t.Errorf("HashBytes with empty input returned negative: %d", h)
	}
}

func TestBytesToColorDeterministic(t *testing.T) {
	palette := newTestPalette()
	input := []byte("consistent-bytes")
	c1 := BytesToColor(palette, bytes.NewReader(input))
	c2 := BytesToColor(palette, bytes.NewReader(input))
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
		t.Error("BytesToColor is not deterministic")
	}
}

func TestRegularAndBoldColorsDiffer(t *testing.T) {
	pairs := []struct {
		name    string
		regular ColorStringer
		bold    ColorStringer
	}{
		{"Black", Black, BBlack},
		{"Red", Red, BRed},
		{"Green", Green, BGreen},
		{"Yellow", Yellow, BYellow},
	}
	for _, p := range pairs {
		t.Run(p.name, func(t *testing.T) {
			reg := p.regular("x")
			bld := p.bold("x")
			if reg == bld {
				t.Errorf("%s regular and bold produce identical output: %q", p.name, reg)
			}
		})
	}
}

func TestPurpleIsNotBlue(t *testing.T) {
	purple := Purple("x")
	blue := BBlue("x")
	if purple == blue {
		t.Error("Purple should not produce the same output as BBlue")
	}
}

func TestTealIsNotWhite(t *testing.T) {
	teal := Teal("x")
	white := IWhite("x")
	if teal == white {
		t.Error("Teal should not produce the same output as IWhite")
	}
}

func TestOnIPurpleCode(t *testing.T) {
	result := OnIPurple("x")
	// Should use code 105 (high intensity background purple), not 95
	if result == "" {
		t.Fatal("OnIPurple returned empty")
	}
	// Verify it doesn't contain the old broken code
	if result == fmt.Sprintf("\033[10;95m%s\033[0m", "x") {
		t.Error("OnIPurple still uses broken ANSI code 10;95")
	}
}

func TestDifferentInputsDifferentColors(t *testing.T) {
	palette := newTestPalette()
	type rgba struct{ r, g, b, a uint32 }
	colors := make(map[rgba]bool)
	inputs := []string{"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi"}
	for _, s := range inputs {
		c := StringToColor(palette, s)
		r, g, b, a := c.RGBA()
		colors[rgba{r, g, b, a}] = true
	}
	if len(colors) < 2 {
		t.Errorf("expected multiple distinct colors, got %d", len(colors))
	}
}
