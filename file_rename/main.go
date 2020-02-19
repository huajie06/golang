package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	test_path string = "/Users/huajiezhang/go/src/file_rename"
)

func main() {
	renameFile()
	listPath(test_path)
}

func renameFile() {
	err := filepath.Walk(test_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		pttn := "n_008.txt"
		new_name := "n_1118.txt"
		if t, _ := filepath.Match(pttn, info.Name()); t {
			dir, _ := filepath.Split(path)
			newPath := filepath.Join(dir, new_name)

			err := os.Rename(path, newPath)
			if err != nil {
				log.Println(err)
			}

			// fmt.Println(path)
			// fmt.Println(newPath)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func listPath(p string) {
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

}
