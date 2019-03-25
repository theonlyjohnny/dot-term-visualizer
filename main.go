package main

import (
	"os"

	"github.com/theonlyjohnny/dot-term-visualizer/logger"
	"github.com/theonlyjohnny/dot-term-visualizer/parse"
	"github.com/theonlyjohnny/dot-term-visualizer/render"
)

var (
	log = logger.Log
)

func main() {

	graphs := parse.GetGraphsForPaths(os.Args[1:])

	for _, graph := range graphs {
		renderErr := render.Graph(graph)
		if renderErr != nil {
			log.Errorf("Could not render contents of %s: %s", graph, renderErr.Error())
		}
	}
}
