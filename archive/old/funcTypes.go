package archive

import (
	"fmt"
)

type binFunc func(int, int) int // func types

func add(x, y int) int { return x + y }

func (f binFunc) Error() string { return "binFunc error" }

func mainFuncType() {
	var bb binFunc

	fmt.Println(bb.Error())

	var c binFunc // create c as type binFunc, c now is a func and also has a method
	c = add1
	fmt.Println(c(1, 1))
	fmt.Println(c.Error())
}

func add1(a int, b int) int {
	return a + b
}
