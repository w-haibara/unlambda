package cli

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"unlambda"

	"github.com/peterh/liner"
)

const history historyFileName = ".w-haibara_unlambda_history"

func Run() int {
	run()
	return 0
}

func run() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	history.readHistory(line)
	defer history.writeHistory(line)

	out := &bytes.Buffer{}
	env := unlambda.Env{
		In:  os.Stdin,
		Out: out,
		Err: os.Stderr,
	}

	for {
		expr, err := line.Prompt(">>> ")

		ctx, cancel := context.WithTimeout(context.Background(),
			time.Duration(time.Millisecond*500))
		defer cancel()

		if err == nil {
			var res string
			res, err := eval(env, expr, ctx)
			if err != nil {
				fmt.Errorf(err.Error())
			}

			fmt.Println("")
			fmt.Println("--- Result ---")
			fmt.Println(res)
			fmt.Println("-- Output ---")
			fmt.Println(out)

			line.AppendHistory(expr)
			out.Reset()
		} else if err == liner.ErrPromptAborted {
			cancel()
			break
		} else {
			log.Print("Error reading line: ", err)
			break
		}
	}
}

func eval(env unlambda.Env, expr string, ctx context.Context) (string, error) {
	res, err := env.EvalString(expr, ctx)
	if err != nil {
		return "", err
	}

	return res, nil
}
