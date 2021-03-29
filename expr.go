package unlambda

import (
	"fmt"
	"io"
	"strings"
)

type Expr string

func ToExpr(str string) Expr {
	return Expr(strings.TrimSpace(str))
}

func (e Expr) Tokenize() Token {
	t := Token{}

	for i := 0; i < len(e); i++ {
		v := string(e[i])
		if i >= 1 && v == "." {
			i++
			v += string(e[i])
		}
		t = append(t, v)
	}

	return t
}

func (e Expr) Sprint() string {
	return fmt.Sprintf("expr: [%v]", e)
}

func (e Expr) Fprint(out io.Writer) {
	fmt.Fprintln(out, e.Sprint())
}
