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
	if cursor == 0 {
		e.text = append([]rune{r}, e.text...)
	} else if cursor < len(e.text) {
		e.text = append(e.text[:cursor], e.text[cursor-1:]...)
		e.text[cursor] = r
	} else {
		e.text = append(e.text[:cursor], r)
	}
	if r == rune('\n') {
		e.x = 1
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
	x, y := e.calcCursorXY(e.calcCursor())
	str := fmt.Sprintf("x=%d, y=%d, cursor=%d, len(text)=%d, x,y=%d,%d", e.x, e.y, e.calcCursor(), len(e.text), x, y)
	for i, r := range []rune(str) {
		termbox.SetCell(i, e.height-1, r, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (e *Editor) MoveCursor(x, y int) {
	c := e.calcCursor()

	if x > 0 {
		if c+x <= len(e.text) {
			e.x, e.y = e.calcCursorXY(c + x)
		}
	} else {
		if 0 <= c+x {
			if e.text[c+x] == rune('\n') {
				e.x, e.y = e.calcCursorXY(c + x - 1)
			} else {
				e.x, e.y = e.calcCursorXY(c + x)
			}
		}
	}
}

// CalcCursor calc index of []rune from e.x and e.y.
func (e *Editor) calcCursor() int {
	ri := 0
	y := 1
	x := 0

	for y < e.y {
		for _, r := range e.text {
			ri++
			if r == '\n' {
				y++
				break
			}
		}
	}

	for _, r := range e.text[ri:] {
		if x >= e.x-runewidth.RuneWidth(r) {
			break
		}
		x += runewidth.RuneWidth(r)
		ri++
	}

	return ri
}

// calcCursorXY calc x and y from index of []rume
func (e *Editor) calcCursorXY(index int) (int, int) {
	x := 1
	y := 1
	for i := 0; i < index; i++ {
		if e.text[i] == rune('\n') {
			x = 1
			y++
		} else {
			x = x + runewidth.RuneWidth(e.text[i])
		}
	}
	return x, y
}
