package main

import (
	"image"

	"github.com/nsf/termbox-go"
)

type Cell struct {
	Point   image.Point
	HasMine bool
	Opened  bool
	Flagged bool

	NumMineNeighbor int
}

func (c *Cell) GetTermboxCell() *termbox.Cell {
	if c.Flagged {
		return &termbox.Cell{'F', termbox.ColorRed | termbox.AttrBold, termbox.ColorBlue}
	}

	if !c.Opened {
		return &termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlue}
	}

	if c.HasMine {
		return &termbox.Cell{'*', termbox.ColorWhite, termbox.ColorRed}
	}

	if c.NumMineNeighbor == 0 {
		return &termbox.Cell{' ', termbox.ColorDefault, termbox.ColorDefault}
	}

	return &termbox.Cell{rune('0' + c.NumMineNeighbor), termbox.ColorBlue, termbox.ColorDefault}
}
