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

	flags.IntVar(&fieldOpts.Bomb, "bomb", 5, "Number of Bomb")
	flags.IntVar(&fieldOpts.Width, "width", 5, "Field width")
	flags.IntVar(&fieldOpts.Height, "height", 5, "Field height")

	flags.BoolVar(&version, "version", false, "display the version")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagsError
	}

	// Version
	if version {
		fmt.Fprintf(cli.errStream, "%s v%s\n", Name, Version)
		return ExitCodeOK
	}

	return ExitCodeOK
}

const usage = `
Usage: %s [options]

  MineSweeper

Options:

  -bomb=<num>         Number of bomb
  -width=<width>      Width of a Field
  -height=<height>    Height of a Field

  -version            Print the version of this application
`
