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
		}
	}
}

func Test_add(t *testing.T) {
	testCases := []struct {
		in      node
		inL     node
		inR     node
		inAfter node
		err     error
	}{
		{
			in: node{
				v: "`",
			},
			inL: node{
				v: ".a",
			},
			inR: node{
				v: ".b",
			},
			inAfter: node{
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
			in: node{
				v: "`",
			},
			inL: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: ".b",
				},
			},
			inR: node{
				v: ".c",
			},
			inAfter: node{
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
			in: node{
				v: "`",
			},
			inL: node{
				v: ".a",
			},
			inR: node{
				v: "`",
				l: &node{
					v: ".b",
				},
				r: &node{
					v: ".c",
				},
			},
			inAfter: node{
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
			in: node{
				v: "`",
			},
			inL: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: ".b",
				},
			},
			inR: node{
				v: "`",
				l: &node{
					v: ".c",
				},
				r: &node{
					v: ".d",
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
		err := (&testCase.in).add(testCase.inL, testCase.inR)
		assert.Equal(t, testCase.err, err)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_replaceL(t *testing.T) {
	testCases := []struct {
		in      node
		inL     node
		inAfter node
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
			inL: node{
				v: ".c",
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: ".c",
				},
				r: &node{
					v: ".b",
				},
			},
		},
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
			inL: node{
				v: "`",
				l: &node{
					v: ".c",
				},
				r: &node{
					v: ".d",
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: "`",
					l: &node{
						v: ".c",
					},
					r: &node{
						v: ".d",
					},
				},
				r: &node{
					v: ".b",
				},
			},
		},
	}

	for _, testCase := range testCases {
		(&testCase.in).replaceL(testCase.inL)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_replaceR(t *testing.T) {
	testCases := []struct {
		in      node
		inR     node
		inAfter node
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
			inR: node{
				v: ".c",
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: ".a",
				},
				r: &node{
					v: ".c",
				},
			},
		},
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
			inR: node{
				v: "`",
				l: &node{
					v: ".c",
				},
				r: &node{
					v: ".d",
				},
			},
			inAfter: node{
				v: "`",
				l: &node{
					v: ".a",
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
		},
	}

	for _, testCase := range testCases {
		(&testCase.in).replaceR(testCase.inR)

		if diff := cmp.Diff(testCase.inAfter, testCase.in, cmp.AllowUnexported(testCase.in)); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_String(t *testing.T) {
	testCases := []struct {
		in  node
		out string
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
			out: "`.a.b",
		},
		{
			in: node{
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
			out: "``.a.b.c",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.out, (&testCase.in).String())
	}
}
