package utils

import (
	"math"
	"strconv"
	"strings"
)

func inchesToColumns(inches float64) float64 {
	return inches * colPerInch
}
func inchesToRows(inches float64) float64 {
	return inches * rowPerInch
}

func columnsToInches(columns float64) float64 {
	return columns / colPerInch
}

func rowsToInches(rows float64) float64 {
	return rows / rowPerInch
}

func ptToInch(pt float64) float64 {
	return pt * ptPerInch
}

func ptToColumns(input float64) float64 {
	inches := ptToInch(input)
	return inchesToColumns(inches)
}

func ptToRows(input float64) float64 {
	inches := ptToInch(input)
	return inchesToRows(inches)
}

func getFloatsFromCommaString(v string) []float64 {
	split := strings.Split(strings.Trim(v, "\""), ",")
	value := make([]float64, len(split))
	for i, str := range split {
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
	return ptRect{
		floats[0],
		floats[1],
		floats[2],
		floats[3],
	}
}

func floatToRoundedInt(i float64) int {
	return int(math.Round(i))
}
