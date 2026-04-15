package colorhash

import (
	"encoding/binary"
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
	hashb := h.Sum(nil)
	hashb = hashb[len(hashb)-8:]
	lsb := binary.BigEndian.Uint64(hashb)
	sint := int(lsb)
	if sint < 0 {
		sint = sint + MaxInt
	}
	return sint
}

// HashBytes returns a deterministic non-negative integer hash of the
// data read from r using FNV-64.
func HashBytes(r io.Reader) int {
	h := fnv.New64()
	io.Copy(h, r)
	hashb := h.Sum(nil)
	hashb = hashb[len(hashb)-8:]
	lsb := binary.BigEndian.Uint64(hashb)
	sint := int(lsb)
	if sint < 0 {
		sint = sint + MaxInt
	}
	return sint
}

// BytesToColor hashes the data from r and maps it to a color in p.
func BytesToColor(p ColorSet, r io.Reader) color.Color {
	i := HashBytes(r) % p.Len()
	return p.Get(i)
}

// StringToColor hashes s and maps it to a color in p.
func StringToColor(p ColorSet, s string) color.Color {
	i := HashString(s) % p.Len()
	return p.Get(i)
}

// GetString hashes s and returns it wrapped in the corresponding
// palette entry's ANSI escape codes.
func (sp StringerPalette) GetString(s string) string {
	h := HashString(s)
	h = h % len(sp)
	return sp[h](s)
}
