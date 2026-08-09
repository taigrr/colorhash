package colorhash

import (
	"hash/fnv"
	"image/color"
	"io"
	"math/bits"
)

const (
	// MaxUint is the maximum value of an unsigned integer.
	MaxUint = ^uint(0)
	// MaxInt is the maximum value of a signed integer.
	MaxInt = int(MaxUint >> 1)
	// intSignBit is the bit position that becomes int's sign bit once
	// a uint64 hash sum is truncated to int, on both 32- and 64-bit
	// platforms (bits.UintSize is 32 or 64).
	intSignBit = uint64(1) << (bits.UintSize - 1)
)

// HashString returns a deterministic non-negative integer hash of s
// using FNV-64.
func HashString(s string) int {
	h := fnv.New64()
	io.WriteString(h, s)
	return hashSumToInt(h.Sum64())
}

// HashBytes returns a deterministic non-negative integer hash of the
// data read from r using FNV-64. Read errors from r are ignored (the
// hash is computed over whatever was read before the error); use
// HashReader to detect them.
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

// hashSumToInt maps sum onto a non-negative int by clearing the bit
// that becomes int's sign bit once truncated to the platform's int
// width. On every platform this discards one bit of hash entropy
// (the sign bit); on 32-bit platforms the upper 32 bits of sum are
// additionally discarded by the int conversion itself, independent
// of this masking.
func hashSumToInt(sum uint64) int {
	return int(sum &^ intSignBit)
}

// BytesToColor hashes the data from r and maps it to a color in p.
// It returns nil if p is empty.
func BytesToColor(p ColorSet, r io.Reader) color.Color {
	if p.Len() == 0 {
		return nil
	}
	i := HashBytes(r) % p.Len()
	return p.Get(i)
}

// StringToColor hashes s and maps it to a color in p. It returns nil
// if p is empty.
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
