package archive

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func mainp1() {
	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv")
	csvR := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)

	s, _ := csvR.ReadAll()

	fmt.Printf("Game Start!!\n-----\n")
	for k, v := range s {
		fmt.Printf("Question #%v, %v = ", k+1, v[0])
		scanner.Scan()
		ans := scanner.Text()
		if ans != v[1] {
			fmt.Printf("Game over, you scored %v out of %v.\n", k, len(s))
			break
		}
	}

}
