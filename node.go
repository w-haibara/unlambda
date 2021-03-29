package unlambda

import (
	"fmt"
	"io"
	//	"strings"
)

type Node struct {
	Val    string
	Parent *Node
	Lhs    *Node
	Rhs    *Node
}

func (n Node) IsRoot() bool {
	return n.Parent == nil
}

func (n Node) IsLeaf() bool {
	return n.Lhs == nil && n.Rhs == nil
}

func (n *Node) CheckBranches() bool {
	if n.IsLeaf() {
		return true
	}

	if n.Lhs.Parent != n || n.Rhs.Parent != n {
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

/*
func (n Node) sprintTree(i int) string {
	str := ""
	if !n.IsLeaf() {
		str += strings.Repeat("| ", i) + n.Lhs.Sprint() + "\n"
		str += n.Lhs.sprintTree(i + 1)
		str += strings.Repeat("| ", i) + n.Rhs.Sprint() + "\n"
		str += n.Rhs.sprintTree(i + 1)
	}
	return str
}

func (n *Node) SprintTree() string {
	str := ""
	str += n.SprintFn() + "\n"
	return str + n.sprintTree(1)
}

func (n Node) FprintTree(out io.Writer) {
	fmt.Fprintln(out, "--- tree ---")
	fmt.Fprintln(out, n.SprintTree())
	fmt.Fprintln(out, "---      ---")
}
*/
