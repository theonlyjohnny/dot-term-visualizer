package visualizer

import (
	"errors"
	"os"
	"os/signal"

	"github.com/gdamore/tcell"
)

var (
	fakeExtraRunes = make([]rune, 0) // this is a fallback thing for characters >1 column (some East Asian chars) so they won't be supported for now
)

//I wish this was in render/screen_managment.go but if I want it to be a domain type it can't be... :/

func containsSig(list []tcell.Key, is tcell.Key) bool {
	present := false
	for _, key := range list {
		if is == key {
			present = true
			break
		}
	}
	return present
}

//RenderPoints is a util function to render all the passed Point instances onto the contained Screen
func (s *Screen) RenderPoints(points []Point) {
	for _, point := range points {
		style := tcell.StyleDefault
		if point.Style != nil {
			style = *point.Style
		}
		s.SetContent(
			point.X, point.Y,
			point.Char,
			fakeExtraRunes,
			style,
		)
	}
}

//WaitForExit spawns several goroutines which send an error into the Screen's exitChan when there is a user interaction indicating exit, or if there is a tcell Error
func (s *Screen) WaitForExit() {
	done := false

	eventChan := make(chan tcell.Event)
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	exitKeys := []tcell.Key{tcell.KeyCtrlZ, tcell.KeyCtrlD, tcell.KeyCtrlC}
	go func() {
		for {
			event := s.PollEvent()
			eventChan <- event
		}
	}()

	var err error

	for !done {
		select {
		case event := <-eventChan:
			switch ev := event.(type) {
			case *tcell.EventKey:
				key := ev.Key()
				if containsSig(exitKeys, key) {
					done = true
					continue
				}
				if key == tcell.KeyRune && ev.Rune() == 'q' {
					done = true
					continue
				}
			case *tcell.EventError: // quit
				err = errors.New(ev.Error())
				done = true
				continue
			}
		case _ = <-sigChan:
			done = true
			continue
		}
	}

	s.ExitChan <- err
}
