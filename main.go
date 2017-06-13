package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	e := NewEditor()
	e.SetSize(termbox.Size())
	e.Draw()

	mainLoop(e)
}

func mainLoop(e *Editor) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowLeft, termbox.KeyCtrlB:
				e.MoveCursor(-1, 0)
				e.Draw()
			case termbox.KeyArrowRight, termbox.KeyCtrlF:
				e.MoveCursor(1, 0)
				e.Draw()
			case termbox.KeyBackspace, termbox.KeyBackspace2:
			case termbox.KeyDelete, termbox.KeyCtrlD:
			case termbox.KeyTab:
			case termbox.KeySpace:
			case termbox.KeyCtrlK:
			case termbox.KeyHome, termbox.KeyCtrlA:
			case termbox.KeyEnd, termbox.KeyCtrlE:
			case termbox.KeyEnter:
				e.AddRune(rune('\n'))
				e.Draw()
			default:
				if ev.Ch != 0 {
					e.AddRune(ev.Ch)
					e.Draw()
				}
			}
		case termbox.EventResize:
			e.SetSize(termbox.Size())
		}
	}
}
