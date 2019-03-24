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
	inches := columns / colPerInch
	return inches
}

func rowsToInches(rows float64) float64 {
	inches := rows / rowPerInch
	return inches
}

func ptToInch(pt float64) float64 {
	inches := pt * ptPerInch
	return inches
}

func ptToColumns(input float64) float64 {
	inches := ptToInch(input)
	columns := inchesToColumns(inches)
	return columns
}

func ptToRows(input float64) float64 {
	inches := ptToInch(input)
	rows := inchesToRows(inches)
	return rows
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

func round(i float64) int {
	return int(math.Round(i))
}
