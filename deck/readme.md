## stringer


### code steps

- need to run `go get golang.org/x/tools/cmd/stringer` to get stringer
- run `go generate`
- run `go test`


### TODO

- all above
- iota
  - Go's iota identifier is used in const declarations to simplify
    definitions of incrementing numbers. Because it can be used in
    expressions, it provides a generality beyond that of simple
    enumerations.
  - Use with `const` normally


### example

```go
package main

import (
	"fmt"
)

type Suit uint8 
// Suit is THE type, jsut value will be uint8
// so a1 != b since they are different type

func main() {

	var a Suit
	var a1 Suit
	var b uint8

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(a1 == a)
	fmt.Println(a1 == b)
}
// it gives error 
// invalid operation: a1 == b (mismatched types Suit and uint8)
```




```go
package main

import (
	"fmt"
)

type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}

func main() {
	fmt.Println(North)
}
// this will print out `North`
// but if remove the `String` func, it will be evaluted as 0
```
