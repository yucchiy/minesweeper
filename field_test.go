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
			if field.Grid[y][x].HasMine {
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

func TestInField_withOutOfRangePosition(t *testing.T) {
	field := &Field{Width: 10, Height: 10}

	if ret := field.InField(-1, -1); ret == true {
		t.Fatalf("expected %q to eq false", ret)
	}

	if ret := field.InField(-1, 10); ret == true {
		t.Fatalf("expected %q to eq false", ret)
	}

	if ret := field.InField(10, -1); ret == true {
		t.Fatalf("expected %q to eq false", ret)
	}

	if ret := field.InField(10, 10); ret == true {
		t.Fatalf("expected %q to eq false", ret)
	}
}

func TestInField_withInField(t *testing.T) {
	field := &Field{Width: 10, Height: 10}

	if ret := field.InField(0, 0); ret == false {
		t.Fatalf("expected %q to eq true", ret)
	}

	if ret := field.InField(0, 9); ret == false {
		t.Fatalf("expected %q to eq true", ret)
	}

	if ret := field.InField(9, 0); ret == false {
		t.Fatalf("expected %q to eq true", ret)
	}

	if ret := field.InField(9, 9); ret == false {
		t.Fatalf("expected %q to eq true", ret)
	}
}

func TestGetNeighbors_withOutOfField(t *testing.T) {
	field := &Field{Width: 2, Height: 2}

	if _, err := field.GetNeighbors(2, 2); err == nil {
		t.Fatalf("expected %q to not eq nil", err)
	}
}

// Todo: add strict testing
func TestGetNeighbors_LeftTop(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(0, 0)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 3
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}

func TestGetNeighbors_TopCenter(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(1, 0)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 5
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}

func TestGetNeighbors_RightTop(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(2, 0)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 3
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}

func TestGetNeighbors_LeftBottom(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(0, 2)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 3
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}

func TestGetNeighbors_RightBottom(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(2, 2)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 3
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}

func TestGetNeighbors_Center(t *testing.T) {
	field := &Field{Width: 3, Height: 3}

	neighbors, err := field.GetNeighbors(1, 1)
	if err != nil {
		t.Fatalf("expected %q to eq nil", err)
	}

	expected := 8
	if ret := len(neighbors); ret != expected {
		t.Fatalf("expected %d to eq %d", ret, expected)
	}
}
