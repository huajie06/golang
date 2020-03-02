package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type Mode int

const (
	ModeCombo Mode = iota
	ModeTriangle
	ModeRect
	ModeEllipse
	ModeCircle
	ModeRotatedRect
	ModeBeziers
	ModeRotatedEllipse
	ModePolygon
)

func main() {
	err := primitive("1.jpeg", 10, withOpt(ModePolygon))
	if err != nil {
		fmt.Println(err)
	}
}

// mode = Polygon, which is 8

func withOpt(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

func primitive(infile string, numShapes int, opts ...func() []string) error {
	outfile := fmt.Sprintf("%s%s", "out_", infile)
	var args []string
	for _, opt := range opts {
		args = append(args, opt()...)
	}

	cmdStr := strings.Fields(fmt.Sprintf("-i %s -o %s -n %d", infile, outfile, numShapes))
	cmdStr = append(cmdStr, args...)

	fmt.Println("primitive", cmdStr)
	cmd := exec.Command("primitive", cmdStr...)
	b, err := cmd.CombinedOutput()

	fmt.Println(string(b))

	if err != nil {
		return err
	}

	return nil
}
