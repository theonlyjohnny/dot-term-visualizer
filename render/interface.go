package render

import (
	"github.com/theonlyjohnny/dot-term-visualizer/render/shapes"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

func getRectFromCommaString(v string) (shapes.Rect, *visualizer.ScaleFactor) {
	ptsRect := shapes.CommaStringToPtRect(v)
	rect := ptsRect.ConvertToRect()
	scaleFactor := rect.GetScaleFactor()
	rect.Scale(scaleFactor)
	return rect, scaleFactor
}

//GetNodeDetails returns row, column, width, height for the given Graphviz "pos", "width", and "height" attributes from a NodeStmt
// func GetNodeDetails(heightStr, widthStr, posStr string, sf *ScaleFactor) (int, int, int, int) {

// widthF := getFloatsFromCommaString(widthStr)[0]   // in
// heightF := getFloatsFromCommaString(heightStr)[0] // in
// pos := getFloatsFromCommaString(posStr)[:2]       // midpoint, [x,y] in pt

// log.Debugf("GetNodeDetails widthF: %v, heightF: %v, pos: %v", widthF, heightF, pos)

// x := pos[0]
// y := pos[1]

// widthBuffer := inchToPt(widthF) / 2.0
// heightBuffer := inchToPt(heightF) / 2.0

// lx := x - widthBuffer
// ly := y - heightBuffer

// ux := x + widthBuffer
// uy := y + heightBuffer

// shapes.Rect := getTermRectFromPtRect(&ptRect{
// lx,
// ly,
// ux,
// uy,
// })

// scale(&shapes.Rect, sf)

// r := rect(shapes.Rect)

// return r.getRoundedRect()
// }
