package unlambda

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_s(t *testing.T) {
	testCases := []struct {
		in      node
		inAfter node
	}{
		{
			in: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: "`",
						l: &node{
							v: "s",
						},
						r: &node{
							v: ".a",
						},
					},
					r: &node{
						v: ".b",
					},
				},
				r: &node{
					v: ".c",
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: ".a",
					},
					r: &node{
						v: ".c",
					},
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
		},
	}

	env := Env{
		//In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	}

	for _, testCase := range testCases {
		env.s(&testCase.in)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}
