package unlambda

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

func Test_k(t *testing.T) {
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
						v: "k",
					},
					r: &node{
						v: ".a",
					},
				},
				r: &node{
					v: ".b",
				},
			},
			inAfter: node{
				v: ".a",
			},
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: "k",
					},
					r: &node{
						v: "`",
						l: &node{
							v: "k",
						},
						r: &node{
							v: ".a",
						},
					},
				},
				r: &node{
					v: "`",
					l: &node{
						v: "k",
					},
					r: &node{
						v: ".b",
					},
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: "k",
				},
				r: &node{
					v: ".a",
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
		env.k(&testCase.in)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_i(t *testing.T) {
	testCases := []struct {
		in      node
		inAfter node
	}{
		{
			in: node{
				v: "`",
				l: &node{
					v: "i",
				},
				r: &node{
					v: ".a",
				},
			},
			inAfter: node{
				v: ".a",
			},
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: "i",
				},
				r: &node{
					v: "`",
					l: &node{
						v: "k",
					},
					r: &node{
						v: ".a",
					},
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: "k",
				},
				r: &node{
					v: ".a",
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
		env.i(&testCase.in)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_dotX(t *testing.T) {
	testCases := []struct {
		in      node
		inAfter node
		out     string
	}{
		{
			in: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: ".b",
				},
			},
			inAfter: node{
				v: ".b",
			},
			out: "a",
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: "`",
					l: &node{
						v: "k",
					},
					r: &node{
						v: ".b",
					},
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: "k",
				},
				r: &node{
					v: ".b",
				},
			},
			out: "a",
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: ".あ",
				},
				r: &node{
					v: ".b",
				},
			},
			inAfter: node{
				v: ".b",
			},
			out: "あ",
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: ".🤔",
				},
				r: &node{
					v: ".b",
				},
			},
			inAfter: node{
				v: ".b",
			},
			out: "🤔",
		},
		{
			in: node{
				v: "`",
				l: &node{
					v: ".👨‍👨‍👧‍👦",
				},
				r: &node{
					v: ".b",
				},
			},
			inAfter: node{
				v: ".b",
			},
			out: "👨‍👨‍👧‍👦",
		},
	}

	out := &bytes.Buffer{}
	env := Env{
		//In:  os.Stdin,
		Out: out,
		Err: os.Stderr,
	}

	for _, testCase := range testCases {
		env.dotX(&testCase.in)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}

		assert.Equal(t, testCase.out, out.String())
		out.Reset()
	}
}
