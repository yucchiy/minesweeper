package main

import (
	"flag"
	"fmt"
	"io"
)

// exit code
const (
	ExitCodeOK    int = 0
	ExitCodeError     = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeBadArgsError
	ExitCodeRuntimeError
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {

	var version bool
	var fieldOpts FieldOpts
	flags := flag.NewFlagSet("minesweeper", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)
	flags.Usage = func() {
		fmt.Fprintf(cli.errStream, usage, Name)
	}

	flags.IntVar(&fieldOpts.Mine, "mine", 5, "Number of Mine")
	flags.IntVar(&fieldOpts.Width, "width", 5, "Field width")
	flags.IntVar(&fieldOpts.Height, "height", 5, "Field height")

	flags.BoolVar(&version, "version", false, "display the version")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	if fieldOpts.Width <= 0 {
		fmt.Fprintf(cli.errStream, "option -width should be a positive number")
		return ExitCodeBadArgsError
	}

	if fieldOpts.Height <= 0 {
		fmt.Fprintf(cli.errStream, "option -height should be a positive number")
		return ExitCodeBadArgsError
	}

	if fieldOpts.Width*fieldOpts.Height < fieldOpts.Mine {
		fmt.Fprintf(cli.errStream, "option -mine should be smaller than field size")
		return ExitCodeBadArgsError
	}

	// Version
	if version {
		fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
		return ExitCodeOK
	}

	game, err := CreateGame(&fieldOpts)
	if err != nil {
		fmt.Fprintf(cli.errStream, "runtime error")
		return ExitCodeRuntimeError
	}

	if err := game.Play(); err != nil {
		fmt.Fprintf(cli.errStream, "runtime error")
		return ExitCodeRuntimeError
	}

	return ExitCodeOK
}

const usage = `
Usage: %s [options]

  MineSweeper

Options:

  -mine=<num>         Number of mine
  -width=<width>      Width of a Field
  -height=<height>    Height of a Field

  -version            Print the version of this application
`
