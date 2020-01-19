package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	// "strings"
	// "unicode"
)

func main() {

	// f := func(c rune) bool {
	// 	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	// }

	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv")

	r := csv.NewReader(file)

	for {
		s, err := r.Read()
		if err == io.EOF {
			break
		}

		// fmt.Println(s)
		// fmt.Printf("type %T", s[1])
		// fmt.Printf("fist is %v, second is %v\n", s[0], s[1])

		fmt.Println(s[0], s[1])
	}

	// fmt.Printf("type %T, value is %v", s, s)
	// fmt.Println(s)
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }
}
