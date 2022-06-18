package archive

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Readincsv is ...
func Readincsv() {
	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv")
	defer file.Close()
	// DEF: func Open(name string) (*File, error)
	// a file descriptor

	r := csv.NewReader(file)
	// DEF: func NewReader(r io.Reader) *Reader
	// io.Reader is the Reader interface, so the underlying data can be anything implemented the Read() method which *File is
	// return Reader is a csv interface, which has a Read() method

	for {
		record, err := r.Read()
		// DEF: func (r *Reader) Read() (record []string, err error)
		if err == io.EOF {
			fmt.Println("reached EOF")
			break
		}
		fmt.Println(record)
	}

}
