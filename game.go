package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"image"
)

type Game struct {
	Field         *Field
	State         GameState
	Cursor        image.Point
	PositionMine  image.Point
	PositionField image.Point

	ChangedField bool
}

func CreateGame(opts *FieldOpts) (*Game, error) {
	field, err := CreateField(opts)
	if err != nil {
		return nil, err
	}

	return &Game{
		Field:         field,
		PositionField: image.Point{X: 1, Y: 2},
		PositionMine:  image.Point{X: field.Width / 2, Y: 1},
		Cursor:        image.Point{X: 1, Y: 2},
		State:         StatePlay}, nil
}

func DrawColorString(str string, x, y int, fg, bg termbox.Attribute) {
	i, j := x, y
	for _, c := range str {
		if c == '\n' {
			i = x
			j += 1
			continue
		}
		termbox.SetCell(i, j, c, fg, bg)
		i += 1
	}
}

func (game *Game) Display(clear bool) {
	if clear {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	}

	DrawColorString("Minesweeper", 0, 0, termbox.AttrBold, termbox.ColorDefault)
	if game.State == StatePlay {
		termbox.SetCursor(game.Cursor.X+1, game.Cursor.Y+1)
	} else {
		termbox.HideCursor()
	}

	if game.ChangedField || clear {
		DrawColorString(fmt.Sprintf("%02d", game.Field.GetNumMineLeft()), game.PositionMine.X, game.PositionMine.Y, termbox.ColorRed|termbox.AttrBold, termbox.ColorWhite)
		DisplayField(game.Field, game.PositionField)
		game.ChangedField = false
	}

	termbox.Flush()
}

func DisplayField(f *Field, pos image.Point) {
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			c := f.Grid[y][x].GetTermboxCell()
			if c != nil {
				termbox.SetCell(x+pos.X, y+pos.Y, c.Ch, c.Fg, c.Bg)
			}
		}
	}
}

func (game *Game) Play() error {

	err := termbox.Init()
	if err != nil {
		return err
	}
	defer termbox.Close()

	clear := true
	game.ChangedField = true
	for game.State = StatePlay; game.State != StateQuit; {
		game.Display(clear)
	}

	return nil
}
