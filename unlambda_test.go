package unlambda

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	in  string
	out string
}

func Test_Render(t *testing.T) {
	testCases := []TestCase{
		// function p
		{
			in:  "`.a.b",
			out: "a",
		},
		{
			in:  "`.ai",
			out: "a",
		},
		{
			in:  "`````````````.H.e.l.l.o.,. .w.o.r.l.d.!i",
			out: "Hello, world!",
		},
		{
			in:  "`.a`.c.d",
			out: "ca",
		},
		{
			in:  "``.a.b`.c.d",
			out: "acb",
		},
		// function r
		{
			in:  "`ri",
			out: "\n",
		},
		{
			in:  "```.ar.bi",
			out: "a\nb",
		},
		// function k
		{
			in:  "``ki`.ai",
			out: "a",
		},
		{
			in:  "```k.aii",
			out: "a",
		},
		{
			in:  "```k.b`.aii",
			out: "ab",
		},
		{
			in:  "```k.a`k.bi",
			out: "a",
		},
		// function s
		{
			in:  "``s`.ai`.bi",
			out: "ab",
		},
		{
			in:  "```s.a.b.c",
			out: "abc",
		},
		{
			in:  "````skk.ai",
			out: "a",
		},
		{
			in:  "```si`ki.a",
			out: "a",
		},
		{
			in:  "```sii```skk.b",
			out: "b",
		},
		{
			in:  "```sii`.a.b",
			out: "ab",
		},
	}

	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fr := bufio.NewWriter(f)

	for _, testCase := range testCases {
		buffer := &bytes.Buffer{}
		op := Option{
			//In:  os.Stdin,
			Out: buffer,
			Err: fr,
		}

		fmt.Fprintln(op.Err, "----------------------")

		expr := ToExpr(testCase.in)
		expr.Fprint(op.Err)

		fmt.Fprintln(op.Err, "")

		token := expr.Tokenize()

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*1))
		op.Eval(ctx, token)
		cancel()

		fmt.Fprintln(op.Err, "----------------------")

		assert.Equal(t, testCase.out, buffer.String())
	}
}
