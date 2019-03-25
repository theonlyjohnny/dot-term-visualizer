package utils

//ScaleFactor is used internally to the utils.scale method, to scale a rect
type ScaleFactor struct {
	columnScaleFactor float64
	rowScaleFactor    float64
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
