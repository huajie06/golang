package urlshort

import (
	"fmt"
)

// Test func
func Test() string {
	// ways to create types
	var ret string

	var a int
	a = 1
	ret += fmt.Sprintf("var a created with value %v\n", a)
	var b int = 1
	ret += fmt.Sprintf("var b created with value %v\n", b)

	var c [2]int
	c[0] = 1
	c[1] = 1
	ret += fmt.Sprintf("var c created with value %v\n", c)
	var d [2]int = [2]int{1, 2}
	var d1 = [2]int{1, 2}
	ret += fmt.Sprintf("var d created with value %v\n", d)
	ret += fmt.Sprintf("var d1 created with value %v\n", d1)

	var e [2][2]int = [2][2]int{{1, 1}, {1, 2}}
	var e1 = [2][2]int{{1, 1}, {1, 2}}
	ret += fmt.Sprintf("var e created with value %v\n", e)
	ret += fmt.Sprintf("var e1 created with value %v\n", e1)

	var f []int = []int{1, 2}
	var f1 = []int{1, 2}
	ret += fmt.Sprintf("var f created with value %v\n", f)
	ret += fmt.Sprintf("var f1 created with value %v\n", f1)

	var g = make([]int, 1)
	g = append(g, 1)
	ret += fmt.Sprintf("var g created with value %v\n", g)

	var h = map[string]int{
		"a": 1,
		"b": 2,
	}
	var h1 = make(map[string]int)
	h1["a"] = 1
	var h2 = map[string]int{}
	h2["a"] = 1
	ret += fmt.Sprintf("var h created with value %v\n", h)
	ret += fmt.Sprintf("var h1 created with value %v\n", h1)
	ret += fmt.Sprintf("var h2 created with value %v\n", h2)

	type i struct {
		a string
		b int
	}
	var i0 = i{}
	ret += fmt.Sprintf("var i0 created with value %v\n", i0)

	var j = map[string]i{}
	var j1 = make(map[string]i)
	j["a"] = i{"string", 1}
	j1["a"] = i{"string", 1}
	ret += fmt.Sprintf("var j created with value %v\n", j)
	ret += fmt.Sprintf("var j1 created with value %v\n", j1)
	return ret
}

// A ...
type A interface {
	m1() int // this type needs to match the method type
	m2() string
}

// B ...
type B struct {
	a int
	b string
}

func (data B) m1() int {
	return data.a
}
func (data B) m2() string {
	return fmt.Sprintf("method with return non-declared type value is %v", data.b)
}

// Test1 ...
func Test1() {
	var b = B{100, "hello"}
	fmt.Println(b)
	fmt.Println(b.m1())

	var c A
	c = B{100, "hello"}
	fmt.Println(c.m1())

	var d A
	d = B{1, "1"}
	fmt.Println(d.m2())

}
