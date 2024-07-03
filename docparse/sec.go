package docparse

import (
	"bytes"
	"encoding/binary"
)

type PSS struct {
	name      [64]byte
	bsize     uint16
	typ       byte
	flag      byte
	left      uint32
	right     uint32
	child     uint32
	guid      [16]uint16
	userflags uint32
	time      [2]uint64
	sstart    uint32
	size      uint32
	_         uint32
}
type Sector []byte

func (s *Sector) Uint32(bit uint32) uint32 {
	return binary.LittleEndian.Uint32((*s)[bit : bit+4])
}

func (s *Sector) NextSid(size uint32) uint32 {
	return s.Uint32(size - 4)
}

func (s *Sector) MsatValues(size uint32) []uint32 {

	return s.values(size, int(size/4-1))
}

func (s *Sector) AllValues(size uint32) []uint32 {

	return s.values(size, int(size/4))
}

func (s *Sector) values(size uint32, length int) []uint32 {

	var res = make([]uint32, length)

	buf := bytes.NewBuffer((*s))

	binary.Read(buf, binary.LittleEndian, res)

	return res
}

func I32(b []byte) uint32 {

	return binary.LittleEndian.Uint32(b[:])
}

func I16(b []byte) uint16 {

	return binary.LittleEndian.Uint16(b[:])
}

func I8(b byte) uint8 {

	return uint8(b)
}
