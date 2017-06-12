package main

import (
	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Editor struct {
	text   []rune
	cursor int
	width  int
	height int
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) SetSize(w, h int) {
	e.width = w
	e.height = h
}

func (e *Editor) AddRune(r rune) {
	if e.cursor < len(e.text) {
		e.text = append(e.text[:e.cursor], e.text[e.cursor-1:]...)
		e.text[e.cursor] = r
	} else {
		e.text = append(e.text[:e.cursor], r)
	}
	if r == rune('\n') {
		e.cursor++
	} else {
		e.cursor += runewidth.RuneWidth(r)
	}
}

func (e *Editor) DeletePrevRune() {
}

func (e *Editor) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
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
	termbox.SetCursor(e.calcCursorXY())
	termbox.Flush()
}

func (e *Editor) MoveCursor(x, y int) {
	if x > 0 {
		if e.cursor+x < len(e.text)+1 {
			e.cursor += x
		}
	} else {
		if 0 <= e.cursor+x {
			e.cursor += x
		}
	}
}

func (e *Editor) calcCursorXY() (int, int) {
	x := 0
	y := 0
	for _, r := range e.text[:e.cursor] {
		if r == rune('\n') {
			x = 0
			y++
		} else {
			x += runewidth.RuneWidth(r)
		}
	}
	return x, y
}
