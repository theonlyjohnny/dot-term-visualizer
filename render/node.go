package render

// import (
// "github.com/awalterschulze/gographviz/ast"
// "github.com/theonlyjohnny/dot-term-visualizer/visualizer"
// )

// func addNodeStmt(elem visualizer.Element, screen visualizer.Screen, storage visualizer.GraphStorage) {
// var node visualizer.Shape

// nodeStmt := elem.Stmt.(*ast.NodeStmt)
// nodeID := nodeStmt.NodeID.String()
// attrsMap := nodeStmt.Attrs.GetMap()

// heightStr, heightOk := attrsMap["height"]
// widthStr, widthOk := attrsMap["width"]
// posStr, posOk := attrsMap["pos"]

// if heightOk && widthOk && posOk {
// // node := utils.GetNodeDetails(heightStr, widthStr, posStr, storage.ScaleFactor)
// } else {
// log.Warnf("Unknown attrsMap combo for NodeStmt: %s", attrsMap)
// }

// screen.RenderPoints(node.Render())
// }
