package stmt

import (
	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dot-term-visualizer/handlers"
	"github.com/theonlyjohnny/dot-term-visualizer/utils"
)

type sortedStmtList []ast.Stmt

func (l sortedStmtList) Len() int {
	return len(l)
}

func (l sortedStmtList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func getOperationsIndex(e ast.Stmt) (int, func(ast.Stmt, *tview.Grid, *utils.GraphStorage)) {
	switch stmt := e.(type) {
	case ast.GraphAttrs:
		return 0, handlers.AddGraphAttrs
	case *ast.NodeStmt:
		return 1, handlers.AddNodeStmt
	default:
		log.Warnf("Unknown ast.Stmt type: %#v", stmt)
		return -1, func(ast.Stmt, *tview.Grid, *utils.GraphStorage) {}
	}
}

func (l sortedStmtList) Less(i, j int) bool {
	var iIndex, jIndex int

	iIndex, _ = getOperationsIndex(l[i])
	jIndex, _ = getOperationsIndex(l[j])

	return iIndex < jIndex
}
