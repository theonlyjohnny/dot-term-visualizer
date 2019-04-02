package render

import (
	"github.com/gdamore/tcell"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

var (
	fakeRunes = make([]rune, 0)
)

func newScreen(exitChan chan<- error) (*visualizer.Screen, error) {
	tScreen, err := tcell.NewScreen()
	if err != nil {
		log.Errorf("Couldn't create tScreen %s", err.Error())
		return nil, err
	}
	err = tScreen.Init()
	if err != nil {
		log.Errorf("Couldn't init tScreen %s", err.Error())
		return nil, err
	}
	tScreen.HideCursor()
	tScreen.DisableMouse()
	screen := &visualizer.Screen{
		tScreen,
		exitChan,
	}
	go screen.WaitForExit()
	return screen, nil
}
