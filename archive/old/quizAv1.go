package archive

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func mainQuiz() {
	file, _ := os.Open("//Users/huajiezhang/go/src/project1/data/problems.csv")
	csvR := csv.NewReader(file)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Game Start!!\n")

	for i := 0; ; i++ {
		s, err := csvR.Read() //readin one by one
		if err == io.EOF {
			break
		}

		fmt.Printf("%v = ", s[0])
		scanner.Scan()

		ans := scanner.Text()
		if s[1] != ans {
			fmt.Printf("Game over! You completed %v\n", i)
			break
		}
	}
}
