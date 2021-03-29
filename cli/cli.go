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
	op := unlambda.Option{
		In:  c.inReader,
		Err: c.errWriter,
		Out: c.outWriter,
	}

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		expr := strings.TrimSpace(stdin.Text())
		e := unlambda.ToExpr(expr)
		e.Fprint(op.Out)

		fmt.Println("\n=== tokenize ===")
		t := e.Tokenize()
		t.Fprint(op.Out)

		fmt.Println("\n=== eval ===")
		op.Eval(t)
		println()
	}
	return 0
}
