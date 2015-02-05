package main

import (
	"fmt"
	"image"

	"github.com/nsf/termbox-go"
)

type GameState int

const (
	StatePlay GameState = iota
	StateLose
	StateWin
	StateQuit
)

type Game struct {
	Field         *Field
	State         GameState
	Cursor        image.Point
	PositionMine  image.Point
	PositionField image.Point
	ChangedField  bool
}

func CreateGame(opts *FieldOpts) (*Game, error) {
	field, err := CreateField(opts)
	if err != nil {
		return nil, err
	}

	return &Game{
		Field:         field,
		PositionField: image.Point{X: 0, Y: 1},
		PositionMine:  image.Point{X: field.Width / 2, Y: 0},
		Cursor:        image.Point{X: 0, Y: 0},
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

	if game.ChangedField || clear {
		DrawColorString(fmt.Sprintf("%d", game.Field.GetNumMineLeft()), game.PositionMine.X, game.PositionMine.Y, termbox.ColorRed|termbox.AttrBold, termbox.ColorWhite)
		DisplayField(game.Field, game.PositionField)
		game.ChangedField = false
	}

	DisplayState(game.State, 0, game.PositionField.Y+game.Field.Height+1)

	if game.State == StatePlay {
		termbox.SetCursor(game.Cursor.X+game.PositionField.X, game.Cursor.Y+game.PositionField.Y)
	} else {
		termbox.HideCursor()
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

func DisplayState(state GameState, x, y int) {
	switch state {
	case StateWin:
		DrawColorString("You win!!", x, y, termbox.ColorWhite, termbox.ColorDefault)
	case StateLose:
		DrawColorString("You lose...", x, y, termbox.ColorWhite, termbox.ColorDefault)
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
		action := GetAction(game.State, termbox.PollEvent())
		if action != nil {
			state := action(game)
			clear = state != game.State
			game.State = state
		}
	}

	return nil
}

func (game *Game) MoveCursor(dx, dy int) error {
	nx, ny := game.Cursor.X+dx, game.Cursor.Y+dy
	if !game.Field.InField(nx, ny) {
		return ErrOutOfField
	}

	game.Cursor.X, game.Cursor.Y = nx, ny
	return nil
}

func (game *Game) ToggleFlag() (GameState, error) {
	changed, err := game.Field.ToggleFlag(game.Cursor.X, game.Cursor.Y)
	if err != nil {
		return StateQuit, err
	}

	game.ChangedField = changed
	return StatePlay, nil
}

func (game *Game) Reveal() (GameState, error) {
	changed, err := game.Field.Reveal(game.Cursor.X, game.Cursor.Y)
	if err != nil {
		return StateQuit, err
	}

	switch game.Field.State {
	case FieldStateWon:
		return StateWin, nil
	case FieldStateLose:
		return StateLose, nil
	}

	game.ChangedField = changed
	return StatePlay, nil
}
