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

	if v == symbolAplly && len(*t) != 0 {
		n.v = symbolAplly
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
	return n.l == nil || n.r == nil
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

func (n *node) replace(p node) {
	n.v = p.v
	n.l = p.l
	n.r = p.r
}

func (n *node) String() string {
	var ret string

	ret += n.v

	if !n.isLeaf() {
		ret += n.l.String()
		ret += n.r.String()
	}

	return ret
}
