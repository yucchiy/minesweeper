package main

import (
	"image"
)

type Cell struct {
	Point image.Point
	Mine  bool
}

func (c *Cell) HasMine() bool {
	return c.Mine
}
