package main

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestGetPlayAction_withInvalidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'x'}

	ret := GetPlayAction(event)

	if ret != nil {
		t.Fatalf("expected %q to eq nil", ret)
	}
}

func TestGetPlayAction_withValidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'f'}

	ret := GetPlayAction(event)

	if ret == nil {
		t.Fatalf("expected %q to not eq nil", ret)
	}
}

func TestGetWinAction_withInvalidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'f'}

	ret := GetWinAction(event)

	if ret != nil {
		t.Fatalf("expected %q to eq nil", ret)
	}
}

func TestGetWinAction_withValidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'q'}

	ret := GetWinAction(event)

	if ret == nil {
		t.Fatalf("expected %q to not eq nil", ret)
	}
}

func TestGetLoseAction_withInvalidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'f'}

	ret := GetLoseAction(event)

	if ret != nil {
		t.Fatalf("expected %q to eq nil", ret)
	}
}

func TestGetLoseAction_withValidEvent(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'q'}

	ret := GetLoseAction(event)

	if ret == nil {
		t.Fatalf("expected %q to not eq nil", ret)
	}
}

func TestGetAction(t *testing.T) {
	event := termbox.Event{Type: termbox.EventKey, Ch: 'q'}

	ret := GetAction(StatePlay, event)

	if ret == nil {
		t.Fatalf("expected %q to not eq nil", ret)
	}
}
