package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unlambda"
)

type cli struct {
	inReader  io.Reader
	outWriter io.Writer
	errWriter io.Writer
}

func (c *cli) run() int {
	e := unlambda.Env{
		In:  c.inReader,
		Out: c.outWriter,
		Err: c.errWriter,
	}

	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	expr := strings.TrimSpace(stdin.Text())

	e.EvalString(expr)
	println()

	return 0
}
