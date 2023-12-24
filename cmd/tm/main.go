package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/chzyer/readline"
	"github.com/winebarrel/tm"
)

var (
	version string
)

type Options struct {
	Expr string `arg:"" optional:"" help:"Expression to evaluate."`
}

func parseArgs() *Options {
	var CLI struct {
		Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Options
}

func repl() {
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

		rs, err := tm.Eval(line)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(rs)
	}
}

func main() {
	options := parseArgs()

	if options.Expr != "" {
		rs, err := tm.Eval(options.Expr)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(rs)
	} else {
		repl()
	}
}
