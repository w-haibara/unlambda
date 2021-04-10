package unlambda

import (
	"context"
	"errors"
	"fmt"
	"io"
)

type Env struct {
	In  io.Reader
	Out io.Writer
	Err io.Writer
}

var (
	errCanceled error = errors.New("context canceled")
)

func (env Env) EvalString(expr string, ctx context.Context) (string, error) {
	t, err := tokenize(expr)
	if err != nil {
		return "", err
	}

	n, err := t.parse()
	if err != nil {
		return "", err
	}

	res, err := env.eval(n, ctx)
	if err != nil {
		if err == errCanceled {
			return res, nil
		}
		return res, err
	}

	return res, nil
}

func (env Env) eval(n *node, ctx context.Context) (string, error) {
	fmt.Fprintln(env.Err, n)

	select {
	case <-ctx.Done():
		fmt.Fprintln(env.Err, ctx.Err())
		return n.String(), errCanceled
	default:
	}

	if n == nil {
		return "", nil
	}

	if n.isLeaf() {
		return n.String(), nil
	}

	if n.v != symbolAplly {
		return n.String(), nil
	}

	switch string(n.l.v[0]) {
	case symbolAplly:
		switch {
		case !n.l.isLeaf() && string(n.l.l.v[0]) == symbolS:
			if n.l.r.isLeaf() && n.r.isLeaf() {
				return n.String(), nil
			}

			old := n.String()

			if !n.l.r.isLeaf() {
				if res, err := env.eval(n.l.r, ctx); err != nil {
					return res, err
				}
			}

			if !n.r.isLeaf() {
				if res, err := env.eval(n.r, ctx); err != nil {
					return res, err
				}
			}

			if old == n.String() {
				return old, nil
			}

			if res, err := env.eval(n, ctx); err != nil {
				return res, err
			}

			return n.String(), nil
		case !n.l.isLeaf() && !n.l.l.isLeaf() && string(n.l.l.l.v[0]) == symbolS:
			fmt.Fprintln(env.Err, "=== eval parameter 1 of s ===")
			if res, err := env.eval(n.l.l.r, ctx); err != nil {
				return res, err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 2 of s ===")
			if res, err := env.eval(n.l.r, ctx); err != nil {
				return res, err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 3 of s ===")
			if res, err := env.eval(n.r, ctx); err != nil {
				return res, err
			}

			fmt.Fprintln(env.Err, "=== apply s ===")
			fmt.Fprint(env.Err, n, " --> ")
			env.s(n)
			fmt.Fprintln(env.Err, n)
			fmt.Fprintln(env.Err, "=== break s ===")

			if res, err := env.eval(n, ctx); err != nil {
				return res, err
			}

			return n.String(), nil
		case !n.l.isLeaf() && string(n.l.l.v[0]) == symbolK:
			fmt.Fprintln(env.Err, "=== eval parameter 1 of k ===")
			if res, err := env.eval(n.l.r, ctx); err != nil {
				return res, err
			}

			fmt.Fprintln(env.Err, "=== eval parameter 2 of k ===")
			if res, err := env.eval(n.r, ctx); err != nil {
				return res, err
			}

			fmt.Fprintln(env.Err, "=== apply k ===")
			fmt.Fprint(env.Err, n, " --> ")
			env.k(n)
			fmt.Fprintln(env.Err, n)
			fmt.Fprintln(env.Err, "=== break k ===")

			if res, err := env.eval(n, ctx); err != nil {
				return res, err
			}

			return n.String(), nil
		default:
			if res, err := env.eval(n.l, ctx); err != nil {
				return res, err
			}

			if res, err := env.eval(n, ctx); err != nil {
				return res, err
			}

			return n.String(), nil
		}
	case symbolS:
		if n.r.isLeaf() {
			return n.String(), nil
		}

		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolK:
		if n.r.isLeaf() {
			return n.String(), nil
		}

		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolI:
		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		env.i(n)

		if res, err := env.eval(n, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolDotX:
		if res, err := env.eval(n.r, ctx); err != nil {
			fmt.Fprintln(env.Err, "DotX", n)

			return res, err
		}

		env.dotX(n)

		if res, err := env.eval(n, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolR:
		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		env.r(n)

		if res, err := env.eval(n, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolV:
		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		env.v(n)

		if res, err := env.eval(n, ctx); err != nil {
			return res, err
		}

		return n.String(), nil
	case symbolE:
		if res, err := env.eval(n.r, ctx); err != nil {
			return res, err
		}

		env.e(n)

		return n.String(), errCanceled
	default:
		return n.String(), fmt.Errorf("unknown symbol[%s]", string(n.l.v[0]))
	}

	return n.String(), fmt.Errorf("unknown error")
}
