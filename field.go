package main

import (
	"fmt"
	"image"
	"math/rand"
	"time"
)

type FieldOpts struct {
	Mine   int
	Width  int
	Height int
}

type Field struct {
	Mine           int
	NumMineFlagged int
	Width          int
	Height         int
	Grid           [][]Cell
}

var ErrInvalidFieldOpts = fmt.Errorf("invalid FieldOpts")
var ErrNoAllocGrid = fmt.Errorf("should alloc grid")

func CreateField(opts *FieldOpts) (*Field, error) {
	if opts == nil {
		return nil, ErrInvalidFieldOpts
	}

	field := &Field{
		Mine:   opts.Mine,
		Width:  opts.Width,
		Height: opts.Height,
	}

	if err := field.Reset(); err != nil {
		return nil, err
	}

	return field, nil
}

func (field *Field) Reset() error {
	if err := field.AllocGrid(); err != nil {
		return err
	}

	if err := field.FillMines(); err != nil {
		return err
	}

	return nil
}

func (field *Field) AllocGrid() error {
	field.Grid = make([][]Cell, field.Height)
	for y := 0; y < field.Height; y++ {
		field.Grid[y] = make([]Cell, field.Width)
		for x := 0; x < field.Width; x++ {
			field.Grid[y][x].Point = image.Point{X: x, Y: y}
		}
	}

	return nil
}

func (field *Field) FillMines() error {
	if field.Grid == nil || len(field.Grid) == 0 || len(field.Grid[0]) == 0 {
		return ErrNoAllocGrid
	}

	rand.Seed(time.Now().Unix())
	mines := make(map[image.Point]bool)
	for i := 0; i < field.Mine; i++ {
		pos := image.Point{X: rand.Intn(field.Width), Y: rand.Intn(field.Height)}
		for mines[pos] {
			pos.X, pos.Y = rand.Intn(field.Width), rand.Intn(field.Height)
		}

		mines[pos], field.Grid[pos.Y][pos.X].Mine = true, true
	}

	field.Mine = len(mines)
	field.NumMineFlagged = 0

	return nil
}

func (field *Field) GetNumMineLeft() int {
	return field.Mine - field.NumMineFlagged
}
