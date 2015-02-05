package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRun__versionFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("minesweeper -version", " ")

	status := cli.Run(args)
	if status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprintf("minesweeper v%s", Version)
	if !strings.Contains(errStream.String(), expected) {
		t.Errorf("expected %q to eq %q", errStream.String(), expected)
	}
}

func TestRun_parseError(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("minesweeper -difficulty extreme", " ")

	status := cli.Run(args)
	if status != ExitCodeParseFlagsError {
		t.Errorf("expected %d to eq %d", status, ExitCodeParseFlagsError)
	}

	expected := "flag provided but not defined: -difficulty"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %q to contain %q", errStream.String(), expected)
	}
}

func TestRun_BadArgsErrorWidth(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("minesweeper -width=-1", " ")

	status := cli.Run(args)
	if status != ExitCodeBadArgsError {
		t.Errorf("expected %d to eq %d", status, ExitCodeBadArgsError)
	}

	expected := "option -width should be a positive number"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %q to contain %q", errStream.String(), expected)
	}
}

func TestRun_BadArgsErrorHeight(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("minesweeper -height=-1", " ")

	status := cli.Run(args)
	if status != ExitCodeBadArgsError {
		t.Errorf("expected %d to eq %d", status, ExitCodeBadArgsError)
	}

	expected := "option -height should be a positive number"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %q to contain %q", errStream.String(), expected)
	}
}

func TestRun_BadArgsErrorMine(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("minesweeper -width=10 -height=10 -mine=101", " ")

	status := cli.Run(args)
	if status != ExitCodeBadArgsError {
		t.Errorf("expected %d to eq %d", status, ExitCodeBadArgsError)
	}

	expected := "option -mine should be smaller than field size"
	if !strings.Contains(errStream.String(), expected) {
		t.Fatalf("expected %q to contain %q", errStream.String(), expected)
	}
}
