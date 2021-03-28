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
		Out: c.outWriter,
		F:   unlambda.DefaultFn,
	}

	for {
		fmt.Print("> ")

		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		expr := strings.TrimSpace(stdin.Text())
		e := unlambda.ToExpr(expr)
		e.Print()

		fmt.Println("\n=== tokenize ===")
		t := e.Tokenize()
		t.Print()

		fmt.Println("\n=== eval ===")
		op.Eval(t)
		println()
	}
	return 0
}
