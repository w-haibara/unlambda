package unlambda

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	in  string
	out string
}

func Test_Render(t *testing.T) {
	testCases := []TestCase{
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
			in:  "`ri",
			out: "\n",
		},
		{
			in:  "```.ar.bi",
			out: "a\nb",
		},
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
			F:   DefaultFn,
		}

		expr := ToExpr(testCase.in)
		token := expr.Tokenize()
		op.Eval(token)

		assert.Equal(t, testCase.out, buffer.String())
	}
}
