package unlambda

import ()

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

	return ret, nil
}

func (t *token) consume() string {
	ret := (*t)[0]
	*t = (*t)[1:]
	return ret
}
