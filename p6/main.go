package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "saveChangesInTheEditorAbcdefg"
	fmt.Println(camelcase2(x))

	// y := "middle-Outz"
	y := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(caesarCipher(y, 2))
}

func camelcase2(s string) int32 {
	var c int32 = 1
	min, max := 'A', 'Z'
	for _, v := range s {
		if v >= min && v <= max {
			c++
		}
	}
	return c
}

func camelcase(s string) int32 {
	var c int32 = 1
	for _, v := range s {
		if strings.ToUpper(string(v)) == string(v) {
			c++
		}
	}
	return c
}

func caesarCipher(s string, k int32) string {
	var ret string

	var _index int
	var _tmp string
	var _char string

	strLower := "abcdefghijklmnopqrstuvwxyz"
	strUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	mL := make(map[string]int)
	mU := make(map[string]int)

	for i, v := range strLower {
		mL[string(v)] = i
	}
	for i, v := range strUpper {
		mU[string(v)] = i
	}

	for _, v := range s {
		_char = string(v)
		_, okL := mL[_char]
		_, okU := mU[_char]
		if !okL && !okU {
			_tmp = _char

		} else if strings.ToLower(_char) == _char {
			_index, _ = mL[string(v)]
			_tmp = string(strLower[(int32(_index)+k)%int32(len(strLower))])

		} else if strings.ToUpper(_char) == _char {
			_index, _ = mU[string(v)]
			_tmp = string(strUpper[(int32(_index)+k)%int32(len(strUpper))])
		}

		ret += _tmp
	}

	return ret
}
