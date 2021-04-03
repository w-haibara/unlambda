package unlambda

import (
	"fmt"
)

type node struct {
	v string
	l *node
	r *node
}

func (t *token) parse() (n *node, err error) {
	v := t.consume()
	n = new(node)

	if v == "`" && len(*t) != 0 {
		n.v = "`"
		n.l, err = t.parse()
		if err != nil {
			return nil, err
		}

		n.r, err = t.parse()
		if err != nil {
			return nil, err
		}
	}

	n.v = v

	return n, nil
}

func (n node) isLeaf() bool {
	return n.l == nil && n.r == nil
}

func (n *node) add(l, r node) error {
	if !n.isLeaf() {
		return fmt.Errorf("node is not leaf")
	}

	n.l = &l
	n.r = &r

	return nil
}

func (n *node) replaceL(l node) {
	n.l = &l
}

func (n *node) replaceR(r node) {
	n.r = &r
}
