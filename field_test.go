package main

import (
	"testing"
)

func TestCreateField_withValidOpts(t *testing.T) {
	opts := &FieldOpts{
		Width:  10,
		Height: 10,
		Bomb:   10,
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

func TestInit(t *testing.T) {
	field := &Field{Width: 10, Height: 8}

	field.Init()

	if len(field.Grid) != 8 {
		t.Fatalf("length of field.Grid should be 8")
	}

	for y := 0; y < 8; y++ {
		if len(field.Grid[y]) != 10 {
			t.Fatalf("length of each field.Grid column should be 10")
		}
	}
}
