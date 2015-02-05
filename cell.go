package main

import (
	"image"

	"github.com/nsf/termbox-go"
)

type Cell struct {
	Point   image.Point
	Mine    bool
	Opened  bool
	Flagged bool

	NumMineNeighbor int
}

func (c *Cell) HasMine() bool {
	return c.Mine
}

func (c *Cell) HasOpened() bool {
	return c.Opened
}

func (c *Cell) HasFlagged() bool {
	return c.Flagged
}

func (c *Cell) GetTermboxCell() *termbox.Cell {
	if c.HasFlagged() {
		return &termbox.Cell{'F', termbox.ColorRed | termbox.AttrBold, termbox.ColorBlue}
	}

	if !c.HasOpened() {
		return &termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlue}
	}

	if c.HasMine() {
		return &termbox.Cell{'*', termbox.ColorWhite, termbox.ColorRed}
	}

	if c.NumMineNeighbor == 0 {
		return &termbox.Cell{' ', termbox.ColorDefault, termbox.ColorDefault}
	}

	return &termbox.Cell{rune('0' + c.NumMineNeighbor), termbox.ColorBlue, termbox.ColorDefault}
}
