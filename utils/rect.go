package utils

import (
	"math"
)

//ScaleFactor is used internally to the utils.scale method, to scale a rect
type ScaleFactor struct {
	columnScaleFactor float64
	rowScaleFactor    float64
}

type rect interface {
	getRect() (float64, float64, float64, float64)
	getRoundedRectSlice() [4]int
}

type ptRect [4]float64
type termRect [4]float64

func (r *termRect) getRect() (float64, float64, float64, float64) {
	return r[0], r[1], r[2], r[3]
}

func (r *termRect) getRoundedRectSlice() [4]int {
	return [4]int{round(r[0]), round(r[1]), round(r[2]), round(r[3])}
}

func (r *ptRect) getRect() (float64, float64, float64, float64) {
	return r[0], r[1], r[2], r[3]
}

func (r *ptRect) getRoundedRectSlice() [4]int {
	return [4]int{round(r[0]), round(r[1]), round(r[2]), round(r[3])}
}

func getTermRectFromPtRect(input *ptRect) termRect {
	llx, lly, urx, ury := input.getRect()

	var width, height float64

	ptWidth := math.Abs(llx - urx)
	ptHeight := math.Abs(lly - ury)

	width = ptToColumns(ptWidth)
	height = ptToRows(ptHeight)

	log.Infof("input: %v, lower: (%v,%v), upper: (%v,%v), ptDims: %vx%v, colRowDims: %vx%v", *input, llx, lly, urx, ury, ptWidth, ptHeight, width, height)

	return termRect{llx, lly, width, height}
}

func getPtRectFromCommaString(v string) ptRect {
	floats := getFloatsFromCommaString(v)
	return ptRect{floats[0], floats[1], floats[2], floats[3]}
}

func getScaleFactor(rect *termRect) ScaleFactor {
	columns := rect[2]
	rows := rect[3]

	columnsIn := columnsToInches(columns)
	rowsIn := rowsToInches(rows)

	preAspectRatio := columnsIn / rowsIn
	if columnsIn > rowsIn {
		//max out columns, then bring rows to match
		columns = maxColumns
		rows = columns / preAspectRatio
	} else {
		//max out rows, then bring columns to match
		rows = maxRows
		columns = preAspectRatio * rows
	}

	// log.Debugf("should be scaling %v to %v (col) and %v to %v (row)", rect[2], columns, rect[3], rows)

	columnScaleFactor := columns / rect[2]
	rowScaleFactor := rows / rect[3]

	return ScaleFactor{
		columnScaleFactor,
		rowScaleFactor,
	}
}

func scale(rect *termRect, factor ScaleFactor) {

	newColumns := rect[2]
	newRows := rect[3]

	origColumns := rect[2]
	origRows := rect[3]

	newColumns = origColumns * factor.columnScaleFactor
	// log.Debugf("to get new columns, multiplying %v * %v = %v", newRows, factor.columnScaleFactor, newColumns)

	newRows = origRows * factor.rowScaleFactor
	// log.Debugf("to get new rows, multiplying %v * %v = %v", newColumns, factor.rowScaleFactor, newRows)

	rect[2] = newColumns
	rect[3] = newRows

	// log.Debugf("Scaled rect by %v from (%vx%vx%vx%v) to (%vx%vx%vx%v)", factor, rect[0], rect[1], origColumns, origRows, rect[0], rect[1], newColumns, newRows)
}
