package parse

import (
	"github.com/theonlyjohnny/dot-term-visualizer/logger"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

var (
	log = logger.Log
)

//GetGraphsForPaths returns a slice of ast.Graphs -- the serialized version of dotfiles -- from the paths to dotfiles passed in
func GetGraphsForPaths(paths []string) []*visualizer.Graph {
	graphVizGraphs := getGraphVizGraphs(paths)

	output := make([]*visualizer.Graph, len(graphVizGraphs))
	for i, graph := range graphVizGraphs {
		output[i] = convertGraphVizGraphToVisualizerGraph(graph)
	}
	return output
}
