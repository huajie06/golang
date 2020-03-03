package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
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
	mux := http.NewServeMux()
	mux.Handle("/", renderImage())

	ec := make(chan error, 1)
	mux.Handle("/upload", upload(ec))
	mux.Handle("/display", display(ec))

	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))

	fmt.Println("server started on :8000")
	log.Println(http.ListenAndServe(":8000", mux))
}

func renderImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html := `<html><body>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="image">
				<input id="numberOfShapes" name="numberOfShapes" type="text" value="10">
				<button type="submit">Upload Image</button>
			</form>
			</body></html>`
		fmt.Fprint(w, html)
		return
	}
}

func upload(e chan<- error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}
		defer f.Close()

		ns, err := strconv.Atoi(r.FormValue("numberOfShapes"))
		if err != nil {
			http.Error(w, "something went wrong", http.StatusBadRequest)
			return
		}

		// fmt.Println(header)

		filename := header.Filename
		fmt.Println(filename)

		err = os.Remove("./img/new.jpeg")

		nf, err := os.Create("./img/old.jpeg")
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		_, err = io.Copy(nf, f)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/display", http.StatusFound)
		// it needs to be http.StatusFound, otherwise it would not redirect correctly

		go func(e chan<- error) {
			err = primitive("./img/old.jpeg", "./img/new.jpeg", ns, withOpt(1))
			e <- err
		}(e)
	}
}

func display(e <-chan error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		select {
		case x, ok := <-e:
			if ok {
				fmt.Printf("Value %v was read.\n", x)
			} else {
				fmt.Println("Channel closed!")
			}
		default:
			fmt.Println("No value ready, moving on.")
		}

		html := `<html><body>
<h2>input imgs </h2><br>
<img src="/img/old.jpeg" style="height:30%">
<br>
<br>
<h2>output imgs </h2><br>
<img src="/img/new.jpeg" style="height:30%">
</body></html>`
		fmt.Fprint(w, html)
		return
	}
}

// mode = Polygon, which is 8

func withOpt(mode Mode) func() []string {
	return func() []string {
		return []string{"-m", fmt.Sprintf("%d", mode)}
	}
}

func primitive(infile, outfile string, numShapes int, opts ...func() []string) error {
	// outfile := fmt.Sprintf("%s%s", "out_", infile)
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
