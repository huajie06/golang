package cmd

import (
	"log"
	"os"
)

const fp = "task.txt"

func db(b []byte, act string) error {
	f, err := os.Open(fp)
	if err != nil {
		log.Println(err)
	}

	_, err := f.Write(b)
	if err != nil {
		log.Println(err)
	}

}
