package cli

import "os"

func Run() int {
	return (&cli{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
		inReader:  os.Stdin,
	}).run()
}