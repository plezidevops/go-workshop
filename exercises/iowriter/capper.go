// Implement io.writer to print Hello World to HELLO WORLD

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Capper struct {
	text io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {

	v := []byte(strings.ToUpper(string(p)))
	return c.text.Write(v)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello world")
}

/*

type Writer interface {
	Write(p, []byte) (n int, err error)
}
*/
