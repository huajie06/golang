package main

import "fmt"

func main() {
	var v1 string = "hello world"
	var str1 []byte = []byte("hello world")

	fmt.Println(v1)

	fmt.Println(str1)

	// loop over some strings, strings is slice of byte

	for i, v := range v1 {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	fmt.Println("===========")

	for i, v := range str1 {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}
	fmt.Println("===========")

	//========================= section ==========================
	// only these many types
	// uint8/16/32
	// int8/16/32
	// float32/float64
	var varF32 float32 = 1.2392
	fmt.Println(varF32)

}
