package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func csvparse(s [][]string) []problem {
	// input is 2d slice with string as element type, return is slice with problem as data type
	ret := make([]problem, len(s))
	for i, line := range s {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	file, err := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv")
	if err == nil {
		csvr := csv.NewReader(file)
		scanner := bufio.NewScanner(os.Stdin) // this return a *Scanner type has methods .Scan()

		s, _ := csvr.ReadAll()
		pset := csvparse(s)

		c := make(chan string)

		nt := time.NewTimer(10 * time.Second)
		counter := 0
	loopproblem:
		for k, v := range pset {
			fmt.Printf("Problem #%v, %v = ", k+1, v.q)

			go func() {
				var ans string
				scanner.Scan()       // .Scan() method will take the []byte and return token
				ans = scanner.Text() // .Text() simply return the token in string
				c <- ans
			}()

			select {
			case <-nt.C:
				break loopproblem
			case ans := <-c:
				if ans != v.a {
					break loopproblem
				}
				counter++
			}
		}
		fmt.Printf("!!!GAME OVER!!!\nYou scored %v out of %v\n", counter, len(pset))

	} else {
		fmt.Println(err)
	}

}
