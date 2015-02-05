package main

import (
	"os"
)

const Name string = "minesweeper"
const Version string = "0.1.0"

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
