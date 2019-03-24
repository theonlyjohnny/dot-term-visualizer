package utils

import (
	"math"
	"strconv"
	"strings"
)

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
	return [4]int{int(math.Round(r[0])), int(math.Round(r[1])), int(math.Round(r[2])), int(math.Round(r[3]))}
}

func (r *ptRect) getRect() (float64, float64, float64, float64) {
	return r[0], r[1], r[2], r[3]
}

func (r *ptRect) getRoundedRectSlice() [4]int {
	return [4]int{int(math.Round(r[0])), int(math.Round(r[1])), int(math.Round(r[2])), int(math.Round(r[3]))}
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
	split := strings.Split(strings.Trim(v, "\""), ",")
	var value ptRect
	for i, str := range split[:4] {
		var err error
		var float float64
		float, err = strconv.ParseFloat(str, 64)
		value[i] = float
		if err != nil {
			log.Errorf("couldnt parse int out of comma seperated string: value: %s error: %s", str, err.Error())
		}
	}

	return value
}

func scaleUp(rect *termRect) {
	columns := rect[2]
	rows := rect[3]

	// oldColumns := rect[2]
	// oldRows := rect[3]

	columnsIn := columnsToInches(columns)
	rowsIn := rowsToInches(rows)

	// log.Debugf("columns: %v, rows: %v, columnsIn: %v, rowsIn: %v", columns, rows, columnsIn, rowsIn)
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

	overscaled := (columns > maxColumns) || (rows > maxRows)

	// postAspectRatio := columnsToInches(columns) / rowsToInches(rows)

	// var logPrefix string
	// var logFunc func(msg string, args ...interface{})

	if overscaled {
		// logFunc = log.Warnf
		// logPrefix = "[Overscaled] "
	} else {
		// logFunc = log.Debugf
		rect[2] = columns
		rect[3] = rows
	}
	// logFunc("%sScaled up from %vx%v to %vx%v (%v) vs (%v)", logPrefix, oldColumns, oldRows, columns, rows, preAspectRatio, postAspectRatio)
}
