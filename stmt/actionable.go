package stmt

import (
	"sort"

	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dot-term-visualizer/utils"
)

//ActionableStmt is a helper struct that wraps an ast.Stmt, GraphStorage instance, and associated func to add the stmt to a view
type ActionableStmt struct {
	stmt       ast.Stmt
	storage    *utils.GraphStorage
	actionFunc func(ast.Stmt, *tview.Grid, *utils.GraphStorage)
}

//Act is called with a tview.Grid and will cause the associated ast.Stmt to be added to the view however it should be
func (as ActionableStmt) Act(g *tview.Grid) {
	if as.actionFunc != nil {
		as.actionFunc(as.stmt, g, as.storage)
	}
}

//GetOperatableStmts takes a list of ast.Stmts and a GraphStorage instance, and returns a slice of actionableStmts, ordered
//by when they need to be executed (i.e, graph before all, etc.)
func GetOperatableStmts(stmts []ast.Stmt, storage *utils.GraphStorage) []ActionableStmt {

	sort.Sort(sortedStmtList(stmts))

	var result []ActionableStmt
	for _, stmt := range stmts {
		if index, actionFunc := getOperationsIndex(stmt); index != -1 {
			actionable := ActionableStmt{
				stmt:       stmt,
				storage:    storage,
				actionFunc: actionFunc,
			}
			result = append(result, actionable)
		}
	}

	return result
}
