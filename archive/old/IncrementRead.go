package archive

import (
	"fmt"
	"io"
	"os"
)

func mainIntRead() {
	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv") // For read access.
	defer file.Close()
	// DEF: func Open(name string) (*File, error)

	var b = make([]byte, 10)
	var offs int64 = 0

	for {
		p, err := file.ReadAt(b, offs)
		// DEF: func (f *File) ReadAt(b []byte, off int64) (n int, err error)
		if err == io.EOF {
			fmt.Printf("%v", string(b[:p]))
			break
		}
		fmt.Printf("%v", string(b))
		offs = offs + 10

	}

}
