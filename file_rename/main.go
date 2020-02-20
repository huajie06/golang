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
	//renameFile("n_008.txt", "n_1118.txt")
	//renameFile("n_1118.txt", "n_008.txt")
	listPath(test_path, "*")
}

func renameFile(findPattern, newName string) {
	err := filepath.Walk(test_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}

		// pttn := "n_008.txt"
		// new_name := "n_1118.txt"

		if t, _ := filepath.Match(findPattern, info.Name()); t && !(info.IsDir()) {
			dir, _ := filepath.Split(path)
			newPath := filepath.Join(dir, newName)

			err := os.Rename(path, newPath)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("rename file succesfully")
			// fmt.Println(path)
			// fmt.Println(newPath)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func listPath(p, pattern string) {

	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}

		if t, _ := filepath.Match(pattern, info.Name()); t && !(info.IsDir()) {
			fmt.Println(path)
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}
}
