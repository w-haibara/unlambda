package unlambda

import (
	"fmt"
	"io"
	"strings"
)

type Option struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (op Option) S(n *Node) *Node {
	return n
}

func (op Option) K(n *Node) *Node {
	return n
}

func (op Option) I(n *Node) *Node {
	fmt.Fprintln(op.Err, "<i>", n.Val)

	if n.Parent.IsRoot() {
		n.Parent.Rhs.Parent = nil
		return n.Parent.Rhs
	}

	n.Parent.Parent.Lhs = n.Parent.Rhs
	n.Parent.Rhs.Parent = n.Parent.Parent

	return n.Parent.Rhs
}

func (op Option) P(n *Node) *Node {
	fmt.Fprintln(op.Err, "<p>", n.Val)
	fmt.Fprint(op.Out, string(n.Val[1]))
	return op.I(n)
}

func (op Option) R(n *Node) *Node {
	fmt.Fprintln(op.Err, "<r>")
	fmt.Fprintln(op.Out, "")
	return op.I(n)
}

func (op Option) Eval(t Token) {
	var n *Node = &Node{}
	t.ToNode(n)
	n1 := op.eval(n, 0)

	fmt.Fprintln(op.Err, "\n--- result ---")
	if n1 != nil {
		fmt.Fprintln(op.Err, "IsRoot:", n1.IsRoot())
		n1.FprintTreeFromRoot(op.Err)
	} else {
		fmt.Fprintln(op.Err, "node: nil")
	}

	return
}

func (op Option) eval(n *Node, i int) *Node {
	ind := strings.Repeat(" ", i)

	if n == nil {
		fmt.Fprintln(op.Err, ind, "node: nil")
		return nil
	}

	fmt.Fprintln(op.Err, ind, n.SprintTree())
	fmt.Fprintln(op.Err, ind, n.SprintTreeFromRoot())
	fmt.Fprintln(op.Err, ind, "node:", n.Sprint())
	fmt.Fprintln(op.Err, ind, "func:", n.SprintFn())

	if !n.CheckBranches() {
		fmt.Fprintln(op.Err, ind, "branches are invalid")
	}

	if n.IsRoot() && n.IsLeaf() {
		fmt.Fprintln(op.Err, ind, "the node is single")
		return n
	}

	if n.IsRoot() {
		fmt.Fprintln(op.Err, ind, "the node is root")
	}

	if n.IsLeaf() {
		fmt.Fprintln(op.Err, ind, "the node is leaf")
	}

	fmt.Fprintln(op.Err, ind, "val:", n.Val)
	switch string(n.Val[0]) {
	case "`":
		if n.IsLeaf() {
			panic("parameter not found")
		}
		n = op.eval(n.Lhs, i+1)
	case "i":
		n = op.I(n)
		if !n.IsRoot() {
			n = op.eval(n.Parent, i+1)
		}
	case ".":
		n = op.P(n)
		if !n.IsRoot() {
			n = op.eval(n.Parent, i+1)
		}
	case "r":
		n = op.R(n)
		if !n.IsRoot() {
			n = op.eval(n.Parent, i+1)
		}
	default:
		panic(fmt.Sprintln("unknown token:", n.Val))
	}

	return n
}
