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

func TestHasOpened(t *testing.T) {
	cell := &Cell{Opened: true}

	expected := true
	if ret := cell.HasOpened(); ret != expected {
		t.Errorf("expected %q to eq %q", ret, expected)
	}
}

func TestHasFlagged(t *testing.T) {
	cell := &Cell{Flagged: true}

	expected := true
	if ret := cell.HasFlagged(); ret != expected {
		t.Errorf("expected %q to eq %q", ret, expected)
	}
}

func TestGetTermboxCell(t *testing.T) {
	cell := &Cell{Flagged: true}

	if ret := cell.GetTermboxCell(); ret == nil {
		t.Errorf("expected %q to not eq nil", ret)
	}
}
