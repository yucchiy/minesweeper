package main

import (
	"github.com/nsf/termbox-go"
)

type Action func(g *Game) GameState

func GetAction(state GameState, event termbox.Event) Action {
	// Todo: しゅっとしたい
	switch state {
	case StatePlay:
		return GetPlayAction(event)
	case StateWin:
		return GetWinAction(event)
	case StateLose:
		return GetLoseAction(event)
	default:
	}
	return nil
}

func GetPlayAction(event termbox.Event) Action {
	switch event {
	case termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowUp}:
		return ActionMoveUp
	case termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowDown}:
		return ActionMoveDown
	case termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowLeft}:
		return ActionMoveLeft
	case termbox.Event{Type: termbox.EventKey, Key: termbox.KeyArrowRight}:
		return ActionMoveRight
	case termbox.Event{Type: termbox.EventKey, Ch: 'h'}:
		return ActionMoveLeft
	case termbox.Event{Type: termbox.EventKey, Ch: 'j'}:
		return ActionMoveDown
	case termbox.Event{Type: termbox.EventKey, Ch: 'k'}:
		return ActionMoveUp
	case termbox.Event{Type: termbox.EventKey, Ch: 'l'}:
		return ActionMoveRight
	case termbox.Event{Type: termbox.EventKey, Ch: 'f'}:
		return ActionToggleFlag
	case termbox.Event{Type: termbox.EventKey, Ch: 'q'}:
		return ActionQuit
	case termbox.Event{Type: termbox.EventKey, Key: termbox.KeySpace}:
		return ActionReveal
	}

	return nil
}

func GetLoseAction(event termbox.Event) Action {
	switch event {
	case termbox.Event{Type: termbox.EventKey, Ch: 'q'}:
		return ActionQuit
	}

	return nil
}

func GetWinAction(event termbox.Event) Action {
	switch event {
	case termbox.Event{Type: termbox.EventKey, Ch: 'q'}:
		return ActionQuit
	}

	return nil
}

func ActionMoveUp(g *Game) GameState {
	g.MoveCursor(0, -1)
	return StatePlay
}

func ActionMoveDown(g *Game) GameState {
	g.MoveCursor(0, 1)
	return StatePlay
}

func ActionMoveLeft(g *Game) GameState {
	g.MoveCursor(-1, 0)
	return StatePlay
}

func ActionMoveRight(g *Game) GameState {
	g.MoveCursor(1, 0)
	return StatePlay
}

func ActionQuit(g *Game) GameState {
	return StateQuit
}

func ActionToggleFlag(g *Game) GameState {
	// Todo: ちゃんとエラー処理
	state, _ := g.ToggleFlag()
	return state
}

func ActionReveal(g *Game) GameState {
	state, err := g.Reveal()
	// Todo: ちゃんとエラー処理
	if err != nil {
		return StateQuit
	}

	return state
}
