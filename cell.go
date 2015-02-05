package main

import (
	"github.com/nsf/termbox-go"
	"image"
)

type Cell struct {
	Point   image.Point
	Mine    bool
	Opened  bool
	Flagged bool
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

	if !c.HasOpened() {
		return &termbox.Cell{' ', termbox.ColorWhite, termbox.ColorBlue}
	}

	if c.HasMine() {
		return &termbox.Cell{'*', termbox.ColorWhite, termbox.ColorRed}
	}

	if c.HasFlagged() {
		return &termbox.Cell{'F', termbox.ColorRed | termbox.AttrBold, termbox.ColorBlue}
	}

	return nil
}
