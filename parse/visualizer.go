package parse

import (
	"sort"

	"github.com/awalterschulze/gographviz/ast"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"

	"github.com/theonlyjohnny/dot-term-visualizer/render"
)

var (
	fakeActionFunc = func(ast.Stmt, *visualizer.Screen, *visualizer.GraphStorage) {}
)

func convertGraphVizGraphToVisualizerGraph(input *ast.Graph) *visualizer.Graph {
	graphStorage := &visualizer.GraphStorage{}
	allElements := getElementsFromStmts(input.StmtList, graphStorage)
	graphElement := allElements[0]
	otherElements := allElements[1:]
	return &visualizer.Graph{
		Element:  graphElement,
		Elements: otherElements,
	}
}

func convertGraphVizStmtToVisualizerElement(stmt ast.Stmt, storage *visualizer.GraphStorage) *visualizer.Element {
	var output *visualizer.Element
	if index, actionFunc := getOperationsIndex(stmt); index != -1 {
		output = &visualizer.Element{
			Stmt:       stmt,
			ActionFunc: actionFunc,
			Storage:    storage,
		}
	}
	return output
}

func getOperationsIndex(e ast.Stmt) (int, visualizer.ActionFunc) {
	switch stmt := e.(type) {
	case ast.GraphAttrs:
		return 0, render.AddGraphAttrs
	// case *ast.NodeStmt:
	// return 1, render.AddNodeStmt
	default:
		log.Warnf("Unknown ast.Stmt type: %#v", stmt)
		return -1, fakeActionFunc
	}
}

func getElementsFromStmts(stmts []ast.Stmt, storage *visualizer.GraphStorage) []*visualizer.Element {
	elements := []*visualizer.Element{}
	sort.Sort(sortedStmtList(stmts))
	for _, stmt := range stmts {
		if element := convertGraphVizStmtToVisualizerElement(stmt, storage); element != nil {
			elements = append(elements, element)
		}
	}
	return elements
}
