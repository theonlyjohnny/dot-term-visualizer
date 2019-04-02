package math

import (
	"math"
	"strconv"
	"strings"
)

var (
	//Abs is a passthrough of std::math.Abs
	Abs = math.Abs
)

//CommaStringToFloats parses a comma seperated string into a slice of float64
func CommaStringToFloats(v string) []float64 {
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

//FloatToRoundedInt rounds a float64 into an int
func FloatToRoundedInt(i float64) int {
	return int(math.Round(i))
}

//GetMaxColumns returns the maximum # of columns that can be rendered
func GetMaxColumns() float64 {
	return maxColumns
}

//GetMaxRows returns the maximum # of rows that can be rendered
func GetMaxRows() float64 {
	return maxRows
}

//ColumnsToInches converts columns to inches
func ColumnsToInches(columns float64) float64 {
	return columns / colPerInch
}

//RowsToInches converts rows to inches
func RowsToInches(rows float64) float64 {
	return rows / rowPerInch
}

//PtToColumns converts pts to columns
func PtToColumns(input float64) float64 {
	inches := ptToInch(input)
	return inchesToColumns(inches)
}

//PtToRows converts pts to rows
func PtToRows(input float64) float64 {
	inches := ptToInch(input)
	return inchesToRows(inches)
}

func inchesToColumns(inches float64) float64 {
	return inches * colPerInch
}

func inchesToRows(inches float64) float64 {
	return inches * rowPerInch
}

func inchToPt(inch float64) float64 {
	return inch / ptPerInch
}

func ptToInch(pt float64) float64 {
	return pt * ptPerInch
}
