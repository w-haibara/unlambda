package unlambda

import (
	"fmt"
	"os"
	"strings"
)

var DefaultFn = Fn{
	"i": fnI,
}

var DefaultOption = Option{
	In:  os.Stdin,
	Out: os.Stdout,
	F:   DefaultFn,
}

func fnP(n1, n2 *Node, op Option) *Node {
	fmt.Fprint(op.Out, strings.TrimPrefix(n1.Val, "."))
	return n2
}

func fnI(n1, n2 *Node, op Option) *Node {
	return n2
}
