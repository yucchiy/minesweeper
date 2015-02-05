package main

import (
	"testing"
)

func TestMoveCursor_1(t *testing.T) {
	opts := &FieldOpts{Width: 5, Height: 5, Mine: 5}
	game, _ := CreateGame(opts)

	game.MoveCursor(3, 0)

	expected := 3
	if game.Cursor.X != expected {
		t.Errorf("expected %d to eq %d", game.Cursor.X, expected)
	}

	expected = 0
	if game.Cursor.Y != expected {
		t.Errorf("expected %d to eq %d", game.Cursor.Y, expected)
	}
}
