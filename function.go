package unlambda

import (
	"fmt"
)

const (
	symbolAplly = "`"
	symbolS     = "s"
	symbolK     = "k"
	symbolI     = "i"
	symbolDotX  = "."
	symbolR     = "r"
)

func (env Env) s(n *node) {
	v1 := *n.l.l.r
	v2 := *n.l.r
	v3 := *n.r

	n.l.replaceL(v1)
	n.l.replaceR(v3)

	n.replaceR(node{v: symbolAplly})
	n.r.replaceL(v2)
	n.r.replaceR(v3)
}

func (env Env) k(n *node) {
	n.replace(*n.l.r)
}

func (env Env) i(n *node) {
	n.replace(*n.r)
}

func (env Env) dotX(n *node) {
	str := n.l.v[1:]
	env.i(n)
	fmt.Fprint(env.Out, str)
}

func (env Env) r(n *node) {
	env.i(n)
	fmt.Fprintln(env.Out, "")
}
