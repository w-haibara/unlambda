package unlambda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tokenize(t *testing.T) {
	testCases := []struct {
		in  string
		out token
		err error
	}{
		{
			in:  "`.a.b",
			out: token{"`", ".a", ".b"},
			err: nil,
		},
		{
			in:  "  ` .a       .b  ",
			out: token{"`", ".a", ".b"},
			err: nil,
		},
		{
			in:  "  ` .a       .   ",
			out: token{"`", ".a", ". "},
			err: nil,
		},
		{
			in:  "  ` .あ       .い  ",
			out: token{"`", ".あ", ".い"},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		token, err := tokenize(testCase.in)
		assert.Equal(t, testCase.err, err)
		assert.Equal(t, testCase.out, token)
	}
}

func Test_consume(t *testing.T) {
	testCases := []struct {
		in      token
		inAfter token
		out     string
	}{
		{
			in:      token{"`", ".a", ".b"},
			inAfter: token{".a", ".b"},
			out:     "`",
		},
	}

	for _, testCase := range testCases {
		v := (&testCase.in).consume()
		assert.Equal(t, testCase.out, v)
		assert.Equal(t, testCase.inAfter, testCase.in)
	}
}
