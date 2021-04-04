package cli

import (
	"io"
	"os"
)

func Run() int {
	return (&cli{
		outWriter: os.Stdout,
		errWriter: io.Discard,
		inReader:  os.Stdin,
	}).run()
}
