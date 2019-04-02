package render

import (
	"github.com/theonlyjohnny/dot-term-visualizer/logger"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

var (
	log = logger.Log
)

//Graph takes in a serialized visualizer.Graph and renders it into a *tcell.Screen
func Graph(graph *visualizer.Graph) error {

	exitChan := make(chan error)
	screen, err := newScreen(exitChan)
	if err != nil {
		return err
	}

	graph.Draw(screen)

	for _, element := range graph.Elements {
		element.Draw(screen)
	}

	screen.Show()
	renderErr := <-exitChan
	screen.Fini()
	return renderErr
}
