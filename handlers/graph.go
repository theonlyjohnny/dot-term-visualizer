package handlers

import (
	"github.com/awalterschulze/gographviz/ast"
	"github.com/rivo/tview"
	"github.com/theonlyjohnny/dot-term-visualizer/utils"
)

//AddGraphAttrs adds a bounding box to the grid and generates a ScaleFactor for use with following
//elements
func AddGraphAttrs(stmt ast.Stmt, view *tview.Grid, storage *utils.GraphStorage) {
	var scaleFactor utils.ScaleFactor
	attrsMap := ast.AttrList(stmt.(ast.GraphAttrs)).GetMap()
	for k, v := range attrsMap {
		if k == "bb" {
			var rect [4]int
			rect, scaleFactor = utils.GetRectFromCommaString(v)

			log.Debugf("new Rect: %v", rect)
			log.Debugf("attr def: %v", v)

			view.SetRect(rect[0], rect[1], rect[2], rect[3])
		} else {
			log.Warnf("Unknown key %s for GraphAttrs", k)
		}
	}
	storage.ScaleFactor = &scaleFactor
}
