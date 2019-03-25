package utils

//GetRectFromCommaString returns a [4]int of the pattern [x, y, width, height] and a ScaleFactor for a Graphviz comma-seperated
//"bb" attribute
func GetRectFromCommaString(v string) ([4]int, ScaleFactor) {
	ptsRect := getPtRectFromCommaString(v)
	termRect := getTermRectFromPtRect(&ptsRect)
	scaleFactor := getScaleFactor(&termRect)
	scale(&termRect, scaleFactor)
	return termRect.getRoundedRectSlice(), scaleFactor
}

//GetPosFromNodePosString returns row, column, rowSpan, colSpan for a given Graphviz "pos" attribute on a NodeStmt
func GetPosFromNodePosString(v string) (int, int, int, int) {
	var row, column, rowSpan, colSpan int

	return row, column, rowSpan, colSpan
}
