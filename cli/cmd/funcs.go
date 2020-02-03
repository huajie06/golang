package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type error interface {
	Error() string
}

const (
	fname string = "task.txt"
	sep   string = ","
)

func main() {
	// for i := 1; i <= 3; i++ {
	// 	err := writeTaskToFile("abc hello world", fname)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	var err error
	// var f *os.File
	var r map[int]string

	r, err = getAllTask(fname)
	if err != nil {
		log.Println(err)
	}

	for i := 1; i <= len(r); i++ {
		fmt.Println(i, r[i])
	}

	// err = delNthLine(fname, 2)
	// if err != nil {
	// 	log.Println(err)
	// }

	// r, err = getAllTask(fname)
	// if err != nil {
	// 	log.Println(err)
	// }

	// for i := 1; i <= len(r); i++ {
	// 	fmt.Println(i, r[i])
	// }

}

func writeTaskToFile(s string, fname string) error {
	var f *os.File

	// if _, err := os.Stat(fname); os.IsNotExist(err) {
	// 	f, err = os.Create(fname)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
	// 	f, err = os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	b := []byte(s)
	b = append(b, []byte(sep)[0])

	_, err = f.Write(b)

	if err != nil {
		return err
	}

	return nil

}

func getAllTask(fname string) (map[int]string, error) {
	var ret = map[int]string{}

	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	c := 1
	curr_p := 0

	for i, v := range b {
		if v == []byte(sep)[0] {
			ret[c] = string(b[curr_p:i])
			c++
			curr_p = i + 1
		}
	}
	return ret, nil
}

func delNthLine(fname string, n int) error {
	m, err := getAllTask(fname)
	size := len(m)
	if err != nil {
		return err
	}

	if _, ok := m[n]; ok {
		delete(m, n)
	} else {
		return errors.New("Task id doesn't exsit")
	}

	mds := ""
	for i := 1; i <= size+1; i++ {
		if v, ok := m[i]; ok {
			mds += v + sep
		}
	}

	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	// _, err = f.WriteString(mds)
	_, err = f.Write([]byte(mds))

	if err != nil {
		return err
	}

	return nil
}
