package unlambda

import (
	"fmt"
	"io"
)

type Token []string

func (t Token) Consume() (Token, string) {
	return t[1:], t[0]
}

func (t Token) Sprint() string {
	str := "token: "
	for _, v := range t {
		str += fmt.Sprintf("[%v] ", v)
	}
	return str
}

func (t Token) Fprint(out io.Writer) {
	fmt.Fprintln(out, t.Sprint())
}

func (t Token) SprintStr() string {
	str := "token: "
	for _, v := range t {
		str += string(v)
	}
	return str
}

func (t Token) FprintStr(out io.Writer) {
	fmt.Fprintln(out, t.SprintStr())
}

func (t Token) ToNode(n *Node) int {
	res := 0

	if len(t) == 0 {
		return -1
	}

	var val string
	t, val = t.Consume()
	res++

	if val == "`" {
		n.Val = val
		n.Lhs = &Node{}
		n.Lhs.Parent = n
		n1 := t.ToNode((*n).Lhs)
		res += n1

		n.Rhs = &Node{}
		n.Rhs.Parent = n
		n2 := t[n1:].ToNode((*n).Rhs)
		res += n2

	} else {
		(*n).Val = val
	}

	return res
}
