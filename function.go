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
	n.v = n.l.r.v
	r := n.l.r.r
	n.l = n.l.r.l
	n.r = r
}

func (env Env) i(n *node) {
	n.v = n.r.v
	r := n.r.r
	n.l = n.r.l
	n.r = r
}

func (env Env) p(n *node) {
	fmt.Fprint(env.Out, "a")
}

func (env Env) r(n *node) {
	fmt.Fprintln(env.Out, "")
}
