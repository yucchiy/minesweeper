package main

type GameState int

const (
	StatePlay GameState = iota
	StateLose
	StateWin
	StateQuit
)
