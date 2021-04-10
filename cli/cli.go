package cli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unlambda"
)

func Run() int {
	out := &bytes.Buffer{}
	env := unlambda.Env{
		In:  os.Stdin,
		Out: out,
		Err: os.Stderr,
	}

	fmt.Print("> ")

	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	expr := strings.TrimSpace(stdin.Text())

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(time.Millisecond*500))
	defer cancel()

	go func() {
		sig := <-c
		fmt.Printf("Got %s signal. Aborting...\n", sig)
		cancel()
	}()

	var res string

	defer func() {
		fmt.Println("")
		fmt.Println("--- Result ---")
		fmt.Println(res)
		fmt.Println("-- Output ---")
		fmt.Println(out)
	}()

	res, err := run(env, expr, ctx)
	if err != nil {
		fmt.Errorf(err.Error())
		return 1
	}

	return 0
}

func run(env unlambda.Env, expr string, ctx context.Context) (string, error) {
	res, err := env.EvalString(expr, ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}
