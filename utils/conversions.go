package utils

import "math"

func columnsToInches(input float64) float64 {
	columns := float64(input)
	inches := columns / colPerInch
	return math.Round(inches)
}

func rowsToInches(input float64) float64 {
	rows := float64(input)
	inches := rows / rowPerInch
	return math.Round(inches)
}

func ptToInch(input float64) float64 {
	pt := float64(input)
	inches := pt * ptPerInch
	return inches
}

func ptToColumns(input float64) float64 {
	inches := ptToInch(input)
	columns := inches * colPerInch
	return math.Round(columns)
}

func ptToRows(input float64) float64 {
	inches := ptToInch(input)
	rows := inches * rowPerInch
	return math.Round(rows)
}
