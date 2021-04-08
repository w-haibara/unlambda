package cli

import (
	//	"io"
	"os"
)

func Run() int {
	return (&cli{
		outWriter: os.Stdout,
		errWriter: os.Stdout, //io.Discard,
		inReader:  os.Stdin,
	}).run()
}
