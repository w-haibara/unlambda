package unlambda

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func Test_parser(t *testing.T) {
	testCases := []struct {
		in  token
		out node
		err error
	}{
		{
			in: token{"`", ".a", ".b"},
			out: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: ".b",
				},
			},
			err: nil,
		},
		{
			in: token{"`", "`", ".a", ".b", ".c"},
			out: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: ".a",
					},
					r: &node{
						v: ".b",
					},
				},
				r: &node{
					v: ".c",
				},
			},
			err: nil,
		},
		{
			in: token{"`", ".a", "`", ".b", ".c"},
			out: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: "`",
					l: &node{
						v: ".b",
					},
					r: &node{
						v: ".c",
					},
				},
			},
			err: nil,
		},
		{
			in: token{"`", "`", ".a", ".b", "`", ".c", ".d"},
			out: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: ".a",
					},
					r: &node{
						v: ".b",
					},
				},
				r: &node{
					v: "`",
					l: &node{
						v: ".c",
					},
					r: &node{
						v: ".d",
					},
				},
			},
			err: nil,
		},
	}

	for _, testCase := range testCases {
		node, err := testCase.in.parse()
		assert.Equal(t, testCase.err, err)

		if diff := cmp.Diff(testCase.out, *node, cmp.AllowUnexported(testCase.out)); diff != "" {
			t.Error(diff)
		} else {
			println(diff)
		}
	}
}
