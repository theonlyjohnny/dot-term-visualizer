package handlers

import (
	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dot-term-visualizer/utils"
)

//AddNodeStmt adds a *ast.NodeStmt to the grid
func AddNodeStmt(stmt ast.Stmt, grid *tview.Grid, storage *utils.GraphStorage) {
	var row, column, rowSpan, colSpan int
	nodeStmt := stmt.(*ast.NodeStmt)
	nodeID := nodeStmt.NodeID.String()
	attrsMap := nodeStmt.Attrs.GetMap()
	log.Debugf("attrs for node: %s", attrsMap)

	node := tview.NewBox().SetBorder(true).SetTitle(nodeID)
	log.Debugf("Adding %q node to grid @ (%d,%d)[%dx%d]", nodeID, row, column, rowSpan, colSpan)
	grid = grid.AddItem(node, row, column, rowSpan, colSpan, 0, 0, false)
}
