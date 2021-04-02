package unlambda

import ()

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