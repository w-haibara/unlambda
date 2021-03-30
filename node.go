package unlambda

import (
	"fmt"
	"io"
)

type Node struct {
	Val    string
	Parent *Node
	Lhs    *Node
	Rhs    *Node
}

func (n *Node) Copy(p *Node) *Node {
	n1 := new(Node)

	if !n.IsLeaf() {
		n1.Lhs = n.Lhs.Copy(n1)
		n1.Rhs = n.Rhs.Copy(n1)
	}

	if p != nil {
		n1.Parent = p
		if p.Lhs == n {
			p.Lhs = n1
		} else {
			p.Rhs = n1
		}
	}

	n1.Val = n.Val

	return n1
}

func (n Node) IsRoot() bool {
	return n.Parent == nil
}

func (n Node) IsLeaf() bool {
	return n.Lhs == nil && n.Rhs == nil
}

func (n *Node) IsLhs() bool {
	return !n.IsRoot() && n.Parent.Lhs == n
}

func (n *Node) IsRhs() bool {
	return !n.IsRoot() && n.Parent.Rhs == n
}

func (n *Node) CheckBranches() bool {
	if n.IsRoot() && n.IsLeaf() {
		return true
	}

	if !n.IsRoot() {
		if !(n.Parent.Lhs == n && n.Parent.Rhs != nil) &&
			!(n.Parent.Lhs != nil && n.Parent.Rhs == n) {
			return false
		}

		if n.IsLeaf() {
			return true
		}
	}

	if n.Lhs == nil || n.Rhs == nil {
		return false
	}

	return n.Lhs.CheckBranches() && n.Rhs.CheckBranches()
}

func (n Node) Sprint() string {
	return "[" + n.Val + "]"
}

func (n Node) Fprint(out io.Writer) {
	fmt.Fprintln(out, n.Sprint())
}

func (n Node) sprintTree() string {
	if n.Val != "`" {
		return n.Sprint()
	} else {
		str := "("
		if n.Lhs != nil {
			str += n.Lhs.sprintTree()
		}
		if n.Lhs != nil && n.Rhs != nil {
			str += ", "
		}
		if n.Rhs != nil {
			str += n.Rhs.sprintTree()
		}
		return str + ")"
	}
}

func (n Node) SprintTree() string {
	return "tree: " + n.sprintTree()
}

func (n Node) FprintTree(out io.Writer) {
	fmt.Fprintln(out, n.SprintTree())
}

func (n Node) SprintTreeFromRoot() string {
	for {
		if n.IsRoot() {
			return n.SprintTree()
		}
		n = *n.Parent
	}
	return ""
}

func (n Node) FprintTreeFromRoot(out io.Writer) {
	fmt.Fprintln(out, n.SprintTreeFromRoot())
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

	return fmt.Sprintf("[%v]: [%v] <-- [%v]", v1, v2, v3)
}

func (n Node) FprintFn(out io.Writer) {
	fmt.Fprintln(out, n.SprintFn())
}
