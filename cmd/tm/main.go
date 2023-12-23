package main

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/winebarrel/tm"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{Prompt: "tm> "})

	if err != nil {
		panic(err)
	}

	defer rl.Close()

	for {
		line, err := rl.Readline()

		if err != nil {
			break
		}

		dur, err := tm.Eval(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(dur)
	}
}
