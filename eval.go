package unlambda

import (
	"fmt"
	"io"
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

	err = env.eval(n)
	if err != nil {
		return err
	}

	fmt.Fprint(env.Out, "a")

	return nil
}

func (env Env) eval(n *node) error {

	return nil
}
