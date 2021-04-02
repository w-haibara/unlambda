package unlambda

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Eval(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		// function p
		{
			in:  "`.a.b",
			out: "a",
		},
		/*
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
			{
				in:  "``si`k.a",
				out: "",
			},
		*/
	}

	f, err := os.Create("test.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fr := bufio.NewWriter(f)

	for i, testCase := range testCases {
		buf := new(bytes.Buffer)
		env := Env{
			//In:  os.Stdin,
			Out: buf,
			Err: fr,
		}

		fmt.Fprintln(fr, "--------------------------------------------")
		fmt.Fprintln(fr, i, testCase.in)
		fmt.Println(i, testCase.in)

		env.EvalString(testCase.in)

		fmt.Fprint(fr, "--------------------------------------------\n\n")
		fr.Flush()

		assert.Equal(t, testCase.out, buf.String())
	}
}
