package main

import (
	"fmt"
	"image"
)

type FieldOpts struct {
	Bomb   int
	Width  int
	Height int
}

type Field struct {
	Bomb   int
	Width  int
	Height int
	Grid   [][]Cell
}

var ErrInvalidFieldOpts = fmt.Errorf("invalid FieldOpts")

func CreateField(opts *FieldOpts) (*Field, error) {
	if opts == nil {
		return nil, ErrInvalidFieldOpts
	}

	field := &Field{
		Bomb:   opts.Bomb,
		Width:  opts.Width,
		Height: opts.Height,
	}

	if err := field.Reset(); err != nil {
		return nil, err
	}

	return field, nil
}

func (field *Field) Reset() error {

	field.Grid = make([][]Cell, field.Height)
	for y := 0; y < field.Height; y++ {
		field.Grid[y] = make([]Cell, field.Width)
		for x := 0; x < field.Width; x++ {
			field.Grid[y][x].Point = image.Point{X: x, Y: y}
		}
	}

	return nil
}
