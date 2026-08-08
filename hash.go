package colorhash

import (
	"hash/fnv"
	"image/color"
	"io"
)

const (
	// MaxUint is the maximum value of an unsigned integer.
	MaxUint = ^uint(0)
	// MaxInt is the maximum value of a signed integer.
	MaxInt = int(MaxUint >> 1)
)

// HashString returns a deterministic non-negative integer hash of s
// using FNV-64.
func HashString(s string) int {
	h := fnv.New64()
	io.WriteString(h, s)
	return hashSumToInt(h.Sum64())
}

// HashBytes returns a deterministic non-negative integer hash of the
// data read from r using FNV-64.
func HashBytes(r io.Reader) int {
	sum, _ := HashReader(r)
	return sum
}

// HashReader returns a deterministic non-negative integer hash of the
// data read from r using FNV-64, along with any read error encountered.
func HashReader(r io.Reader) (int, error) {
	h := fnv.New64()
	_, err := io.Copy(h, r)
	return hashSumToInt(h.Sum64()), err
}

func hashSumToInt(sum uint64) int {
	sint := int(sum)
	if sint < 0 {
		sint = sint + MaxInt
	}
	return sint
}

// BytesToColor hashes the data from r and maps it to a color in p.
func BytesToColor(p ColorSet, r io.Reader) color.Color {
	if p.Len() == 0 {
		return nil
	}
	i := HashBytes(r) % p.Len()
	return p.Get(i)
}

// StringToColor hashes s and maps it to a color in p.
func StringToColor(p ColorSet, s string) color.Color {
	if p.Len() == 0 {
		return nil
	}
	i := HashString(s) % p.Len()
	return p.Get(i)
}

// GetString hashes s and returns it wrapped in the corresponding
// palette entry's ANSI escape codes.
func (sp StringerPalette) GetString(s string) string {
	if len(sp) == 0 {
		return s
	}
	h := HashString(s)
	h = h % len(sp)
	return sp[h](s)
}
