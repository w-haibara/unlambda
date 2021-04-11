package unlambda

import (
	"fmt"
)

type token []string

func tokenize(expr string) (token, error) {
	var ret token

	for _, v := range expr {
		s := string(v)
		if len(ret) > 0 && ret[len(ret)-1] == symbolDotX {
			ret[len(ret)-1] += s
			continue
		}

		if s == " " {
			continue
		}

		ret = append(ret, s)
	}

	if len(ret) == 0 {
		return token{}, fmt.Errorf("token length is zero")
	}

	return ret, nil
}

func (t *token) consume() string {
	ret := (*t)[0]
	*t = (*t)[1:]
	return ret
}
