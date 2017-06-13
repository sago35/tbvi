package main

import (
	"fmt"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Editor struct {
	text   []rune
	x      int
	y      int
	cursor int
	width  int
	height int
}

func NewEditor() *Editor {
	return &Editor{
		x: 1,
		y: 1,
	}
}

func (e *Editor) SetSize(w, h int) {
	e.width = w
	e.height = h
}

func (e *Editor) AddRune(r rune) {
	cursor := e.calcCursor()
	if cursor < len(e.text) {
		e.text = append(e.text[:cursor], e.text[cursor-1:]...)
		e.text[cursor] = r
	} else {
		e.text = append(e.text[:cursor], r)
	}
	if r == rune('\n') {
		e.x = 0
		e.y += 1
	} else {
		e.x += runewidth.RuneWidth(r)
	}
}

func (e *Editor) DeletePrevRune() {
}

func (e *Editor) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(e.x-1, e.y-1)
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
	e.debugDraw()
	termbox.Flush()
}

func (e *Editor) debugDraw() {
	str := fmt.Sprintf("x=%d, y=%d, cursor=%d, len(text)=%d", e.x, e.y, e.calcCursor(), len(e.text))
	for i, r := range []rune(str) {
		termbox.SetCell(i, e.height-1, r, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (e *Editor) MoveCursor(x, y int) {
	if x > 0 {
		if e.x+x <= len(e.text)+1 {
			e.x += x
		}
	} else {
		if 1 <= e.x+x {
			e.x += x
		}
	}
}

func (e *Editor) calcCursor() int {
	// e.x と e.y から、[]runeとしての位置を割り出す
	ri := 0
	x := 0

	//if e.x != 0 {
	for _, r := range e.text {
		if x >= e.x-runewidth.RuneWidth(r) {
			break
		}
		x += runewidth.RuneWidth(r)
		ri++
	}
	//}

	return ri
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
