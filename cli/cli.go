package cli

import (
	"io"
)

type cli struct {
	inReader  io.Reader
	outWriter io.Writer
	errWriter io.Writer
}

func (c *cli) run() int {
	return 0
}
