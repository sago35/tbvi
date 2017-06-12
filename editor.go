package main

import (
	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Editor struct {
	text    []rune
	cursorx int
	cursory int
	width   int
	height  int
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) SetSize(w, h int) {
	e.width = w
	e.height = h
}

func (e *Editor) AddRune(r rune) {
	e.text = append(e.text, r)
	if r == rune('\n') {
		e.cursorx = 0
		e.cursory++
	} else {
		e.cursorx += runewidth.RuneWidth(r)
	}
}

func (e *Editor) DeletePrevRune() {
}

func (e *Editor) Draw() {
	x := 0
	y := 0
	for i := 0; i < len(e.text); i++ {
		if e.text[i] == rune('\n') {
			x = 0
			y++
		} else {
			if x < e.width {
				termbox.SetCell(x, y, e.text[i], termbox.ColorDefault, termbox.ColorDefault)
			}
			x = x + runewidth.RuneWidth(e.text[i])
		}
	}
}

func (e *Editor) UpdateCursor() {
	termbox.SetCursor(e.cursorx, e.cursory)
}
