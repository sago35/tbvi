package main

import (
	"testing"
)

func TestCalcCursor(t *testing.T) {
	tests := []struct {
		Editor   Editor
		Expected int
	}{
		{Editor: Editor{text: []rune("abcde"), x: 0, y: 0}, Expected: 0},
		{Editor: Editor{text: []rune("abcde"), x: 1, y: 0}, Expected: 0},
		{Editor: Editor{text: []rune("abcde"), x: 2, y: 0}, Expected: 1},
		{Editor: Editor{text: []rune("abcde"), x: 3, y: 0}, Expected: 2},
		{Editor: Editor{text: []rune("abcde"), x: 4, y: 0}, Expected: 3},
		{Editor: Editor{text: []rune("abcde"), x: 5, y: 0}, Expected: 4},

		{Editor: Editor{text: []rune("あいう"), x: 0, y: 0}, Expected: 0},
		{Editor: Editor{text: []rune("あいう"), x: 1, y: 0}, Expected: 0},
		{Editor: Editor{text: []rune("あいう"), x: 2, y: 0}, Expected: 0},
		{Editor: Editor{text: []rune("あいう"), x: 3, y: 0}, Expected: 1},
		{Editor: Editor{text: []rune("あいう"), x: 4, y: 0}, Expected: 1},
		{Editor: Editor{text: []rune("あいう"), x: 5, y: 0}, Expected: 2},
		{Editor: Editor{text: []rune("あいう"), x: 6, y: 0}, Expected: 2},
	}

	for _, test := range tests {
		if e, g := test.Expected, test.Editor.calcCursor(); e != g {
			t.Errorf("in=(%q,%d,%d), expected=%d, got=%d",
				string(test.Editor.text), test.Editor.x, test.Editor.y, e, g)
		}
	}
}
