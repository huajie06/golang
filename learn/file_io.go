package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var a float64 = 77
	fmt.Println(a / 2)
	for i := 1; i < 3; i++ {
		fmt.Println("xxx")
	}

	for i, v := range "hello world" {
		fmt.Printf("i: %d, v: %v\n", i, string(v))
	}

	for i, v := range strings.Fields("what today is today ha ha") {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	//========================= section ==========================
	var fPath string = "./test_file1.txt"

	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	r1 := bufio.NewReader(file)
	linen := 0
	for {
		line, _, err := r1.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		linen += 1
		fmt.Println("=======")
		fmt.Println(linen)
		fmt.Println(line)
	}

	// read all

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range b {
		fmt.Println(string(v))
	}

	fmt.Println(string(b))

	//========================= section ==========================
	strs1 := []byte("abcdeddd")
	fmt.Println(string(strs1))

	strs := []byte{'a', 'b'}
	fmt.Println(strs)
	fmt.Println()
	fmt.Println(string(strs))

	//========================= section ==========================
	someLongStr := `hahahah test new day
blabla
qqq
bbbzz
dek kdw dk w
dkfwld
`
	var io_r io.Reader = strings.NewReader(someLongStr)

	r, e := ioutil.ReadAll(strings.NewReader(someLongStr))
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(r)

	//========================= section ==========================
	b1, e := io.ReadAll(io_r)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(string(b1))

	//========================= section ==========================
	nr := bufio.NewReader(io_r)
	p := make([]byte, 5000)
	n, e := nr.Read(p)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(n)
	fmt.Println(string(p[:n]))

	//========================= section ==========================
	nr1 := bufio.NewReader(io_r)

	for {
		line1, _, e := nr1.ReadLine()
		if e == io.EOF {
			break
		}
		if e != nil {
			panic(e)
		}
		fmt.Println(string(line1))
	}

}
