package unlambda

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
)

type Option struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (op Option) S(ctx context.Context, n *Node) *Node {
	defer func() {
		if !n.CheckBranches() {
			fmt.Fprintln(op.Err, "branches are invalid")
			panic("branches are invalid")
		}
	}()

	fmt.Fprintln(op.Err, "---- s parameter 1 ----")
	n.Parent.Rhs = op.eval(ctx, n.Parent.Rhs.Copy(nil), 0)
	n.Parent.Rhs.Parent = n.Parent
	fmt.Fprintln(op.Err, "---- 1111111111111 ----")

	if !n.Parent.IsLhs() || n.Parent.Parent.Val != "`" {
		return n.Parent
	}

	fmt.Fprintln(op.Err, "---- s parameter 2 ----")
	n.Parent.Parent.Rhs = op.eval(ctx, n.Parent.Parent.Rhs.Copy(nil), 0)
	n.Parent.Parent.Rhs.Parent = n.Parent
	fmt.Fprintln(op.Err, "---- 2222222222222 ----")

	if !n.Parent.Parent.IsLhs() || n.Parent.Parent.Parent.Val != "`" {
		return n.Parent.Parent
	}

	fmt.Fprintln(op.Err, "---- s parameter 3 ----")
	n.Parent.Parent.Parent.Rhs = op.eval(ctx, n.Parent.Parent.Parent.Rhs.Copy(nil), 0)
	n.Parent.Parent.Parent.Rhs.Parent = n.Parent.Parent
	fmt.Fprintln(op.Err, "---- 3333333333333 ----")

	root := n.Parent.Parent.Parent
	n1 := n.Parent.Rhs
	n2 := n.Parent.Parent.Rhs
	n3 := n.Parent.Parent.Parent.Rhs
	n4 := new(Node)
	n4 = n3.Copy(n4)

	root.Lhs.Lhs = n1
	n1.Parent = root.Lhs

	root.Lhs.Rhs = n3
	n3.Parent = root.Lhs

	root.Rhs = &Node{
		Val:    "`",
		Parent: root,
	}

	root.Rhs.Lhs = n2
	n2.Parent = root.Rhs

	root.Rhs.Rhs = n4
	n4.Parent = root.Rhs

	return root
}

func (op Option) K(ctx context.Context, n *Node) *Node {
	defer func() {
		if !n.CheckBranches() {
			fmt.Fprintln(op.Err, "branches are invalid")
			panic("branches are invalid")
		}
	}()

	fmt.Fprintln(op.Err, "---- k parameter 1 ----")
	n.Parent.Rhs = op.eval(ctx, n.Parent.Rhs.Copy(nil), 0)
	n.Parent.Rhs.Parent = n.Parent
	fmt.Fprintln(op.Err, "---- 1111111111111 ----")

	if !n.Parent.IsLhs() || n.Parent.Parent.Val != "`" {
		return n.Parent
	}

	fmt.Fprintln(op.Err, "---- k parameter 2 ----")
	n.Parent.Parent.Rhs = op.eval(ctx, n.Parent.Parent.Rhs.Copy(nil), 0)
	n.Parent.Parent.Rhs.Parent = n.Parent.Parent
	fmt.Fprintln(op.Err, "---- 2222222222222 ----")

	if n.Parent.Parent.IsRoot() {
		n.Parent.Rhs.Parent = nil
		return n.Parent.Rhs
	}

	if n.Parent.Parent.Parent.Lhs == n.Parent.Parent {
		n.Parent.Parent.Parent.Lhs = n.Parent.Rhs
	} else {
		n.Parent.Parent.Parent.Rhs = n.Parent.Rhs
	}

	n.Parent.Rhs.Parent = n.Parent.Parent.Parent

	return n.Parent.Rhs
}

func (op Option) I(ctx context.Context, n *Node) *Node {
	fmt.Fprintln(op.Err, "---- i parameter   ----")
	n.Parent.Rhs = op.eval(ctx, n.Parent.Rhs.Copy(nil), 0)
	n.Parent.Rhs.Parent = n.Parent
	fmt.Fprintln(op.Err, "---- 1111111111111 ----")

	n.Parent.Rhs = op.eval(ctx, n.Parent.Rhs.Copy(nil), 0)
	n.Parent.Rhs.Parent = n.Parent

	if n.Parent.IsRoot() {
		n.Parent.Rhs.Parent = nil
		return n.Parent.Rhs
	}

	n.Parent.Parent.Lhs = n.Parent.Rhs
	n.Parent.Rhs.Parent = n.Parent.Parent

	return n.Parent.Rhs
}

func (op Option) P(ctx context.Context, n *Node) *Node {
	n1 := op.I(ctx, n)
	fmt.Fprint(op.Out, string(n.Val[1]))
	fmt.Fprintln(op.Err, "PRINT:"+string(n.Val[1]))
	return n1
}

func (op Option) R(ctx context.Context, n *Node) *Node {
	n1 := op.I(ctx, n)
	fmt.Fprintln(op.Out, "")
	fmt.Fprintln(op.Err, "PRINT: \\n")
	return n1
}

func (op Option) Eval(ctx context.Context, t Token) {
	var n *Node = &Node{}
	t.ToNode(n)

	n = op.eval(ctx, n, 0)

	fmt.Fprintln(op.Err, "\n--- result ---")
	if n != nil {
		if n.IsRoot() && n.IsLeaf() {
			fmt.Fprintln(op.Err, "the node is single")
		}

		if n.IsRoot() {
			fmt.Fprintln(op.Err, "the node is root")
		}

		if n.IsLeaf() {
			fmt.Fprintln(op.Err, "the node is leaf")
		}

		n.FprintFn(op.Err)

		n.FprintTreeFromRoot(op.Err)
	} else {
		fmt.Fprintln(op.Err, "node: nil")
	}

	return
}

func (op Option) eval(ctx context.Context, n *Node, i int) *Node {
	select {
	case <-ctx.Done():
		fmt.Fprintln(op.Err, ctx.Err())
		os.Exit(1)
	default:
	}

	ind := strings.Repeat(" ", i)

	if n == nil {
		fmt.Fprintln(op.Err, ind, "node: nil")
		return nil
	}

	defer func() {
		if !n.CheckBranches() {
			fmt.Fprintln(op.Err, ind, "branches are invalid")
			panic("branches are invalid")
		}
	}()

	fmt.Fprintln(op.Err, ind, n.SprintTree())
	fmt.Fprintln(op.Err, ind, n.SprintTreeFromRoot())
	fmt.Fprintln(op.Err, ind, "node:", n.Sprint())
	fmt.Fprintln(op.Err, ind, "func:", n.SprintFn())

	if n.IsRoot() {
		fmt.Fprintln(op.Err, ind, "the node is root")
	}

	if n.IsLeaf() {
		fmt.Fprintln(op.Err, ind, "the node is leaf")
	}

	if n.IsRoot() && n.IsLeaf() {
		fmt.Fprintln(op.Err, ind, "the node is single")
		return n
	}

	if n.IsLeaf() && n.Parent.Rhs == n {
		fmt.Fprintln(op.Err, ind, "the node is single parameter")
		return n
	}
	fmt.Fprintln(op.Err, ind, "val:", n.Val)
	switch string(n.Val[0]) {
	case "`":
		fmt.Fprintln(op.Err, "<`>")

		if n.IsLeaf() {
			panic("parameter not found")
		}

		if n.Lhs.Val == "k" && n.Rhs.IsLeaf() && (n.IsRoot() || n.IsRhs() || n.Parent.Val != "`") {
			fmt.Fprintln(op.Err, ind, "the node is function k with 1 parameter")
			return n
		}

		if n.Lhs.Val == "s" && n.Rhs.IsLeaf() && (n.IsRoot() || n.IsRhs() || n.Parent.Val != "`") {
			fmt.Fprintln(op.Err, ind, "the node is function s with 1 parameter")
			return n
		}

		if n.Lhs.Val == "s" && n.Rhs.IsLeaf() && n.Parent.Rhs.IsLeaf() && (n.Parent.IsRoot() || n.Parent.Parent.Val != "`" || n.Parent.IsRhs()) {
			fmt.Fprintln(op.Err, ind, "the node is function s with 2 parameter")
			return n
		}

		n = op.eval(ctx, n.Lhs, i+1)
	case "s":
		if n.Parent.Val != "`" {
			return n
		}

		fmt.Fprintln(op.Err, "<s>")
		n = op.eval(ctx, op.S(ctx, n), i+1)
	case "k":
		if n.Parent.Val != "`" {
			return n
		}

		fmt.Fprintln(op.Err, "<k>")
		n = op.eval(ctx, op.K(ctx, n), i+1)
	case "i":
		if n.Parent.Val != "`" {
			return n
		}

		fmt.Fprintln(op.Err, "<i>")
		n = op.eval(ctx, op.I(ctx, n), i+1)
	case ".":
		if n.Parent.Val != "`" {
			return n
		}

		fmt.Fprintln(op.Err, "<p>")
		n = op.eval(ctx, op.P(ctx, n), i+1)
	case "r":
		if n.Parent.Val != "`" {
			return n
		}

		fmt.Fprintln(op.Err, "<r>")
		n = op.eval(ctx, op.R(ctx, n), i+1)
	default:
		panic(fmt.Sprintln("unknown token:", n.Val))
	}

	return n
}
