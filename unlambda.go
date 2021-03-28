package unlambda

import (
	"fmt"
	"io"
	"strings"
)

type Expr string

func ToExpr(str string) Expr {
	return Expr(strings.TrimSpace(str))
}

func (e Expr) Tokenize() Token {
	t := Token{}

	for i := 0; i < len(e); i++ {
		v := string(e[i])
		if i >= 1 && v == "." {
			i++
			v += string(e[i])
		}
		t = append(t, v)
	}

	return t
}

func (e Expr) Sprint() string {
	return fmt.Sprintf("expr: [%v]", e)
}

func (e Expr) Fprint(out io.Writer) {
	fmt.Fprintln(out, e.Sprint())
}

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

type Node struct {
	Val    string
	Parent *Node
	Lhs    *Node
	Rhs    *Node
}

func (n Node) IsRoot() bool {
	return n.Parent == nil
}

func (n Node) sprint() string {
	if n.Val != "`" {
		return "[" + n.Val + "]"
	} else {
		str := "("
		if n.Lhs != nil {
			str += (*n.Lhs).sprint()
		}
		str += ", "
		if n.Rhs != nil {
			str += (*n.Rhs).sprint()
		}
		return str + ")"
	}
}

func (n Node) Sprint() string {
	return "node: " + n.sprint()
}

func (n Node) Fprint(out io.Writer) {
	fmt.Println(out, n.Sprint())
}

func (n Node) SprintFromRoot() string {
	for {
		if n.Parent == nil {
			return n.Sprint()
		}
		n = *n.Parent
	}
	return ""
}

func (n Node) FprintFromRoot(out io.Writer) {
	fmt.Fprintln(out, n.SprintFromRoot())
}

func (n Node) SprintFn() string {
	var v1, v2, v3 string
	if &n != nil {
		v1 = n.Val
	}
	if n.Lhs != nil {
		v2 = n.Lhs.Val
	}
	if n.Rhs != nil {
		v3 = n.Rhs.Val
	}

	return fmt.Sprintf("[%v]: [%v] --> [%v]", v1, v2, v3)
}

func (n Node) FprintFn(out io.Writer) {
	fmt.Fprintln(out, n.SprintFn())
}

type Fn map[string](func(*Node, *Node, Option) *Node)

type Option struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
	F   Fn
}

func (op Option) Eval(t Token) {
	var n Node
	t.ToNode(&n)
	op.eval(&n)
	return
}

func (op Option) eval(n *Node) *Node {
	if n == nil {
		return nil
	}

	if n.Val == "`" {
		if n.Lhs == nil || n.Rhs == nil {
			return nil
		} else if strings.HasPrefix(n.Lhs.Val, ".") {
			n1 := fnP(n.Lhs, n.Rhs, op)
			if n.IsRoot() {
				return n1
			} else {
				n.Parent.Lhs = n1
				n.Parent.Lhs.Lhs = nil
				n.Parent.Lhs.Rhs = nil
				op.eval(n.Parent)
			}
		} else {
			f, ok := op.F[n.Lhs.Val]
			if !ok {
				n.Lhs = op.eval(n.Lhs)
				n.Rhs = op.eval(n.Rhs)
				return nil
			}
			n1 := f(n.Lhs, n.Rhs, op)
			if n.IsRoot() {
				return n1
			} else {
				n.Parent.Lhs = n1
				n.Parent.Lhs.Lhs = nil
				n.Parent.Lhs.Rhs = nil
				op.eval(n.Parent)
			}
		}
	}
	return n
}
