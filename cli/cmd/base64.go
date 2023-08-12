package cmd

import (
	"fmt"
	"os"
	"strings"
)

// in shell if you do
// echo "a" | base64
// it by default include a newline character \n

type encoding struct {
	encodeMap [64]byte
	decodeMap map[uint]uint
	padChar   rune
}

// bitwise operation demo
// https://play.golang.org/p/VeLCx_4orSW

const mapEncoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func newEnc(encoder string) *encoding {
	e := new(encoding)
	e.padChar = '='
	copy(e.encodeMap[:], encoder) // a good way to assign value
	// when you have a slice make([]int{}, 10), you can dynamically pop the value
	// by using copy(dest, src []Type), slice[:] <=> string
	// dest=[0 0 0 0 0 0 0] src=[1 2 3] => [1 2 3 0 0 0 0]

	e.decodeMap = map[uint]uint{}
	for i, v := range mapEncoder {
		e.decodeMap[uint(v)] = uint(i)
	}
	return e
}

func (enc *encoding) encode(src []byte) []byte {
	if len(src) == 0 {
		return []byte{}
	}

	dst := make([]byte, (len(src)+2)/3*4)

	// 1 byte = 8 bits, so everything 3`letters`/bytes = 24 bits => 4 new `blocks`
	// everything 3 letters => 4 new letters, remaining will fill accordingly
	// so the length will have to be `(n + 2) / 3 * 4`

	di, si := 0, 0
	n := (len(src) / 3) * 3

	for si < n {
		// Convert 3x 8bit source bytes into 4 bytes
		val := uint(src[si+0])<<16 | uint(src[si+1])<<8 | uint(src[si+2])
		// if src[si+0] = "a", "a"
		// 01100001 => 011000010000000000000000, add sixteen zeros
		// 01100001 =>         0110000100000000, add 8 zeros
		// 01100001 =>                 01100001, stays the same
		//          => 011000010110000101100001 => 68371921 in decimal
		// use |(or), to get an overall number

		dst[di+0] = enc.encodeMap[val>>18&0x3F]
		dst[di+1] = enc.encodeMap[val>>12&0x3F]
		dst[di+2] = enc.encodeMap[val>>6&0x3F]
		dst[di+3] = enc.encodeMap[val&0x3F]
		// 0x3f = 63 = 111111
		// 011000                   & 111111 => 011000 => 24 => enc.encode[24]
		// 011000010110             & 111111 => 010110 => 22 =>
		// 011000010110000101       & 111111 => 000101 => 5  =>
		// 011000010110000101100001 & 111111 => 100001 => 33 =>
		si += 3
		di += 4
	}

	remain := len(src) - si
	if remain == 0 {
		return dst
	}

	switch remain {
	case 2:
		val := uint(src[si+0])<<10 | uint(src[si+1])<<2
		dst[di+0] = enc.encodeMap[val>>12&0x3f]
		dst[di+1] = enc.encodeMap[val>>6&0x3F]
		dst[di+2] = enc.encodeMap[val&0x3F]
		dst[di+3] = byte(enc.padChar)
	case 1:
		val := uint(src[si+0]) << 4
		dst[di+0] = enc.encodeMap[val>>6&0x3f]
		dst[di+1] = enc.encodeMap[val&0x3F]
		dst[di+2] = byte(enc.padChar)
		dst[di+3] = byte(enc.padChar)
	}
	return dst
}

func Base64(s []string) {
	e := newEnc(mapEncoder)
	strs := strings.Join(s[1:], " ")

	if s[0] == "-d" {
		r := e.encode([]byte(strs))
		fmt.Println(string(r))
	} else if s[0] == "-e" {
		r := e.decode([]byte(strs))
		fmt.Println(string(r))
	} else {
		fmt.Fprintf(os.Stderr, "flags supported are only -d/-e")
	}
}

// decode is not bullet proof, needs to think of some error handling logic.
func (enc *encoding) decode(src []byte) []byte {
	// 1. trim suffix ==, and count
	// 2. determine the return length, (#n*6 -#=*2) / 8. yq== 2*6-2*2 = 1
	// 3. any 4 letters => 3 letters 4*6/8
	l := len(src)

	if l <= 1 {
		return []byte{}
	}

	npad := 0
	if src[l-1] == '=' {
		npad += 1
	}
	if src[l-2] == '=' {
		npad += 1
	}

	dst := make([]uint8, ((l-npad)*6-(npad*2))/8)
	// dst := make([]byte, ((l-npad)*6-(npad*2))/8)

	si, di := 0, 0
	n := (l - npad) / 4 * 4
	for si < n {
		val := enc.decodeMap[uint(src[si+0])]<<18 | enc.decodeMap[uint(src[si+1])]<<12 | enc.decodeMap[uint(src[si+2])]<<6 | enc.decodeMap[uint(src[si+3])]
		dst[di+0] = uint8(val >> 16 & 0xFF)
		dst[di+1] = uint8(val >> 8 & 0xFF)
		dst[di+2] = uint8(val & 0xFF)
		si += 4
		di += 3
	}
	// YQ==
	// l = 4, npad=2
	// n = 0
	// remain = 2
	remain := l - npad - si

	if remain == 0 {
		return dst
	}

	switch remain {
	case 3:
		val := enc.decodeMap[uint(src[si+0])]<<12 | enc.decodeMap[uint(src[si+1])]<<6 | enc.decodeMap[uint(src[si+2])]
		dst[di+0] = uint8(val >> 10 & 0xFF)
		dst[di+1] = uint8(val >> 2 & 0xFF)
	case 2:
		val := enc.decodeMap[uint(src[si+0])]<<2 | enc.decodeMap[uint(src[si+1])]>>4
		dst[di+0] = uint8(val)
	}
	return dst
}
