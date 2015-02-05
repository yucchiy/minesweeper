package main

import (
	"testing"
)

func TestGetTermboxCell(t *testing.T) {
	cell := &Cell{Flagged: true}

	if ret := cell.GetTermboxCell(); ret == nil {
		t.Errorf("expected %q to not eq nil", ret)
	}
}
