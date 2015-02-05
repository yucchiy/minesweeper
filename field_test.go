package main

import (
	"testing"
)

func TestCreateField_withValidOpts(t *testing.T) {
	opts := &FieldOpts{
		Width:  10,
		Height: 10,
		Mine:   10,
	}

	field, err := CreateField(opts)

	if field == nil {
		t.Fatalf("field should not be nil")
	}

	if err != nil {
		t.Fatalf("err should be nil")
	}
}

func TestCreateField_withNil(t *testing.T) {
	field, err := CreateField(nil)

	if field != nil {
		t.Fatalf("field should be nil")
	}

	if err == nil {
		t.Fatalf("err should not be nil")
	}
}

func TestReset(t *testing.T) {
	field := &Field{Width: 10, Height: 8}

	field.Reset()
}

func TestAllocGrid(t *testing.T) {
	field := &Field{Width: 10, Height: 8}

	field.AllocGrid()

	expected := 8
	if len(field.Grid) != expected {
		t.Fatalf("expected %d to eq %d", len(field.Grid), expected)
	}

	expected = 10
	for y := 0; y < 8; y++ {
		if len(field.Grid[y]) != expected {
			t.Fatalf("expected %d to eq %d", len(field.Grid[y]), expected)
		}
	}
}

func TestFillMines_withNoAllocationGrid(t *testing.T) {
	field := &Field{Width: 10, Height: 8}

	err := field.FillMines()

	if err == nil {
		t.Fatalf("err should not be nil")
	}
}

func TestFillMines_success(t *testing.T) {
	field := &Field{Width: 10, Height: 8, Mine: 5}

	field.AllocGrid()
	if err := field.FillMines(); err != nil {
		t.Fatalf("err should be nil")
	}

	mine := 0
	for y := 0; y < field.Height; y++ {
		for x := 0; x < field.Width; x++ {
			if field.Grid[y][x].HasMine() {
				mine++
			}
		}
	}

	if field.Mine != mine {
		t.Errorf("expected %d to eq %d", field.Mine, mine)
	}
}

func TestGetNumMineLeft(t *testing.T) {
	field := &Field{Mine: 10, NumMineFlagged: 4}

	expected := 6
	if ret := field.GetNumMineLeft(); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}
