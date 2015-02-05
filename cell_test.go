package main

import (
	"testing"
)

func TestHasMine(t *testing.T) {
	cell := &Cell{Mine: true}

	expected := true
	if ret := cell.HasMine(); ret != expected {
		t.Errorf("expected %q to eq %q", ret, expected)
	}
}
