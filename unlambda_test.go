package unlambda

import (
	"bytes"
	"fmt"
	"os"
	"testing"

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
	}

	for _, testCase := range testCases {
		buffer := &bytes.Buffer{}
		op := Option{
			//In:  os.Stdin,
			Out: buffer,
			Err: os.Stderr,
		}

		fmt.Fprintln(op.Err, "----------------------")
		expr := ToExpr(testCase.in)
		expr.Fprint(op.Err)
		fmt.Fprintln(op.Err, "")
		token := expr.Tokenize()
		op.Eval(token)
		fmt.Fprintln(op.Err, "----------------------")

		assert.Equal(t, testCase.out, buffer.String())
	}
}
