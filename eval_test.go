package unlambda

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_eval(t *testing.T) {
	testCases := []struct {
		in      string
		inAfter string
		out     string
		err     error
	}{

		// function p
		{
			in:      "`.a.b",
			inAfter: ".b",
			out:     "a",
			err:     nil,
		},
		{
			in:      "`.ai",
			inAfter: "i",
			out:     "a",
			err:     nil,
		},
		{
			in:      "`````````````.H.e.l.l.o.,. .w.o.r.l.d.!i",
			inAfter: "i",
			out:     "Hello, world!",
			err:     nil,
		},
		{
			in:      "`.a`.c.d",
			inAfter: ".d",
			out:     "ca",
			err:     nil,
		},
		{
			in:      "``.a.b`.c.d",
			inAfter: ".d",
			out:     "acb",
			err:     nil,
		},
		// function r
		{
			in:      "`ri",
			inAfter: "i",
			out:     "\n",
			err:     nil,
		},
		{
			in:      "```.ar.bi",
			inAfter: "i",
			out:     "a\nb",
			err:     nil,
		},
		// function i
		{
			in:      "`i.a",
			inAfter: ".a",
			out:     "",
			err:     nil,
		},
		{
			in:      "``i`ii.a",
			inAfter: ".a",
			out:     "",
			err:     nil,
		},
		// function k
		{
			in:      "``k.x`.a.y",
			inAfter: ".x",
			out:     "a",
			err:     nil,
		},
		{
			in:      "```k.aii",
			inAfter: "i",
			out:     "a",
			err:     nil,
		},
		{
			in:      "```k.b`.aii",
			inAfter: "i",
			out:     "ab",
			err:     nil,
		},
		{
			in:      "```k.a`k.bi",
			inAfter: "i",
			out:     "a",
			err:     nil,
		},
		{
			in:      "`k`ki",
			inAfter: "`k`ki",
			out:     "",
			err:     nil,
		},
		// function s
		{
			in:      "``s`.ai`.bi",
			inAfter: "``sii",
			out:     "ab",
			err:     nil,
		},
		{
			in:      "```s.a.b.c",
			inAfter: ".c",
			out:     "abc",
			err:     nil,
		},
		{
			in:      "````skk.ai",
			inAfter: "i",
			out:     "a",
			err:     nil,
		},
		{
			in:      "```si`ki.a",
			inAfter: "i",
			out:     "a",
			err:     nil,
		},
		{
			in:      "```sii```skk.b",
			inAfter: ".b",
			out:     "b",
			err:     nil,
		},
		{
			in:      "```sii`.a.b",
			inAfter: ".b",
			out:     "ab",
			err:     nil,
		},
		{
			in:      "``si`k.a",
			inAfter: "``si`k.a",
			out:     "",
			err:     nil,
		},
		{
			in:      "```s``si`k.*`ki```s``s`k``s`ksk``sii``si`k``s``s`kski``s``s`ksk``s``s`kski",
			inAfter: "i",
			out:     "************************************************************************************************************************************************************************************************************************",
			err:     nil,
		},
	}

	f, err := os.Create("test.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fr := bufio.NewWriter(f)
	defer fr.Flush()

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

		token, err := tokenize(testCase.in)
		if err != nil {
			panic(err)
		}

		node, err := token.parse()
		if err != nil {
			panic(err)
		}

		err = env.eval(node)

		fmt.Fprint(fr, "--------------------------------------------\n\n")

		assert.Equal(t, testCase.err, err)
		assert.Equal(t, testCase.out, buf.String())
		assert.Equal(t, testCase.inAfter, node.String())
	}
}
