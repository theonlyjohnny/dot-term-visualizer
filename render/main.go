package render

import (
	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dot-term-visualizer/logger"
	"github.com/theonlyjohnny/dot-term-visualizer/stmt"
	"github.com/theonlyjohnny/dot-term-visualizer/utils"
)

var (
	log = logger.Log
)

//Graph takes in a serialized ast.Graph and renders it into a *tview.Grid
func Graph(graph *ast.Graph) error {
	storage := utils.GraphStorage{}
	grid := tview.NewGrid()

	operatable := stmt.GetOperatableStmts(graph.StmtList, &storage)
	log.Debugf("operatable: %#v", operatable)
	for _, stmt := range operatable {
		stmt.Act(grid)
	}
	view := grid.SetBorder(true).SetTitle(graph.ID.String())
	application := tview.NewApplication().SetRoot(view, false)
	log.Debug("made application:", application)
	return nil
	// return application.Run()
}
