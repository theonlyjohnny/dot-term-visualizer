package render

import (
	"github.com/awalterschulze/gographviz/ast"
	"github.com/theonlyjohnny/dot-term-visualizer/render/shapes"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

//AddGraphAttrs is an ActionFunc for ast.GraphAttrs
func AddGraphAttrs(stmt ast.Stmt, screen *visualizer.Screen, storage *visualizer.GraphStorage) {
	var scaleFactor *visualizer.ScaleFactor
	attrsMap := ast.AttrList(stmt.(ast.GraphAttrs)).GetMap()
	for k, v := range attrsMap {
		if k == "bb" {
			var rect shapes.Rect
			rect, scaleFactor = getRectFromCommaString(v)

			log.Infof("adding rect %s", rect)
			screen.RenderPoints(rect.Render())

			// utils.RectHollowWithTextAndStyle(screen,
			// rect[0], rect[1],
			// rect[2], rect[3],
			// storage.Graph.ID.String(),
			// tcell.StyleDefault.Background(tcell.ColorDarkGreen),
			// )

		} else {
			log.Warnf("Unknown key %s for GraphAttrs", k)
		}
	}
	storage.ScaleFactor = scaleFactor
}
