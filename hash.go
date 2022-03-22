package go_colorhash

import (
	"crypto/md5"
	"encoding/binary"
	"io"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func HashString(s string) int {
	h := md5.New()
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
	h := md5.New()
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
