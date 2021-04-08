package unlambda

import (
	"fmt"
	"io"
	"reflect"
)

type Env struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

func (env Env) EvalString(expr string) error {
	t, err := tokenize(expr)
	if err != nil {
		return err
	}

	n, err := t.parse()
	if err != nil {
		return err
	}

	if err := env.eval(n); err != nil {
		return err
	}

	return nil
}

func (env Env) eval(n *node) error {
	fmt.Fprintln(env.Err, n)

	if n == nil {
		return nil
	}

	if n.isLeaf() {
		return nil
	}

	if n.v != symbolAplly {
		return nil
	}

	switch string(n.l.v[0]) {
	case symbolAplly:
		switch {
		case !n.l.isLeaf() && string(n.l.l.v[0]) == symbolS:
			if n.l.r.isLeaf() && n.r.isLeaf() {
				return nil
			}

			old := n

			if !n.l.r.isLeaf() {
				if err := env.eval(n.l.r); err != nil {
					return err
				}
			}

			if !n.r.isLeaf() {
				if err := env.eval(n.r); err != nil {
					return err
				}
			}

			if reflect.DeepEqual(old, n) {
				return nil
			}

			if err := env.eval(n); err != nil {
				return err
			}

			return nil
		case !n.l.isLeaf() && !n.l.l.isLeaf() && string(n.l.l.l.v[0]) == symbolS:
			fmt.Fprintln(env.Err, "=== eval parameter 1 of s ===")
			if err := env.eval(n.l.l.r); err != nil {
				return err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 2 of s ===")
			if err := env.eval(n.l.r); err != nil {
				return err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 3 of s ===")
			if err := env.eval(n.r); err != nil {
				return err
			}

			fmt.Fprintln(env.Err, "=== apply s ===")
			fmt.Fprint(env.Err, n, " --> ")
			env.s(n)
			fmt.Fprintln(env.Err, n)
			fmt.Fprintln(env.Err, "=== break s ===")

			if err := env.eval(n); err != nil {
				return err
			}

			return nil
		case !n.l.isLeaf() && string(n.l.l.v[0]) == symbolK:
			fmt.Fprintln(env.Err, "=== eval parameter 1 of k ===")
			if err := env.eval(n.l.r); err != nil {
				return err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 2 of k ===")
			if err := env.eval(n.r); err != nil {
				return err
			}

			fmt.Fprintln(env.Err, "=== apply k ===")
			fmt.Fprint(env.Err, n, " --> ")
			env.k(n)
			fmt.Fprintln(env.Err, n)
			fmt.Fprintln(env.Err, "=== break k ===")

			if err := env.eval(n); err != nil {
				return err
			}

			return nil
		default:
			if err := env.eval(n.l); err != nil {
				return err
			}

			if err := env.eval(n); err != nil {
				return err
			}

			return nil
		}
	case symbolS:
		if n.r.isLeaf() {
			return nil
		}

		if err := env.eval(n.r); err != nil {
			return err
		}

		return nil
	case symbolK:
		if n.r.isLeaf() {
			return nil
		}

		if err := env.eval(n.r); err != nil {
			return err
		}

		return nil
	case symbolI:
		if err := env.eval(n.r); err != nil {
			return err
		}

		env.i(n)

		if err := env.eval(n); err != nil {
			return err
		}

		return nil
	case symbolDotX:
		if err := env.eval(n.r); err != nil {
			return err
		}

		env.dotX(n)

		if err := env.eval(n); err != nil {
			return err
		}

		return nil
	case symbolR:
		if err := env.eval(n.r); err != nil {
			return err
		}

		env.r(n)

		if err := env.eval(n); err != nil {
			return err
		}

		return nil
	case symbolV:
		if err := env.eval(n.r); err != nil {
			return err
		}

		env.v(n)

		if err := env.eval(n); err != nil {
			return err
		}

		return nil
	default:
		return fmt.Errorf("unknown symbol[%s]", string(n.l.v[0]))
	}

	return fmt.Errorf("unknown error")
}
