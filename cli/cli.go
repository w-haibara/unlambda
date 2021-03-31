package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
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

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
		op.Eval(ctx, t)
		cancel()
		println()
	}
	return 0
}
