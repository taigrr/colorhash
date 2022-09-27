package colorhash

import (
	"encoding/binary"
	"hash/fnv"
	"io"
)

const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

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
