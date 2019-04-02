package shapes

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/theonlyjohnny/dot-term-visualizer/render/math"
	"github.com/theonlyjohnny/dot-term-visualizer/visualizer"
)

type rect [4]float64

func (r *rect) getRect() (float64, float64, float64, float64) {
	return r[0],
		r[1],
		r[2],
		r[3]
}

func (r *rect) set(index int, value float64) {
	r[index] = value
}

func (r *rect) getRoundedRect() (int, int, int, int) {
	return math.FloatToRoundedInt(r[0]),
		math.FloatToRoundedInt(r[1]),
		math.FloatToRoundedInt(r[2]),
		math.FloatToRoundedInt(r[3])
}

//PtRect is a rectangle of shape [llx, lly, upx, upy] using the units of pt (1 inch / 72)
type PtRect struct {
	rect
}

func (r PtRect) String() string {
	llx, lly, upx, upy := r.getRect()

	return fmt.Sprintf("[l:(%v,%v) u:(%vx%v)]", llx, lly, upx, upy)
}

//CommaStringToPtRect returns a PtRect from a comma seperated string (the "bb" graphviz attribute)
func CommaStringToPtRect(v string) PtRect {
	floats := math.CommaStringToFloats(v)
	var subSlice rect
	for i := 0; i <= 3; i++ {
		subSlice[i] = floats[i]
	}
	return PtRect{
		subSlice,
	}
}

//ConvertToRect returns a new Rect from the values in the PtRect
func (r *PtRect) ConvertToRect() Rect {
	llx, lly, urx, ury := r.getRect()

	var width, height float64

	ptWidth := math.Abs(llx - urx)
	ptHeight := math.Abs(lly - ury)

	width = math.PtToColumns(ptWidth)
	height = math.PtToRows(ptHeight)

	newRect := Rect{
		Title:  nil,
		Style:  nil,
		Hollow: true,
		rect: rect{
			llx,
			lly,
			width,
			height,
		},
	}

	log.Infof("PtRect.ConvertToRect input: %s --> %s",
		r,
		newRect,
	)

	return newRect
}

//Rect is a Shape that draws a geometric rectangle (with an optional Title, Hollow option, and Style)
type Rect struct {
	Title  *string
	Style  *tcell.Style
	Hollow bool

	rect
}

func (r Rect) String() string {
	orx, ory, width, height := r.getRect()
	return fmt.Sprintf("[l:(%v,%v)[%vx%v]]", orx, ory, width, height)
}

//Render implements the Shape.Render interface requirement
func (r Rect) Render() []visualizer.Point {
	points := []visualizer.Point{}
	x, y, width, height := r.getRoundedRect()
	for row := y; row < y+height; row++ {
		for column := x; column <= x+width; column++ {
			var char rune
			if column == x {
				if row == y {
					char = tcell.RuneULCorner
				} else if row == height+y {
					char = tcell.RuneLLCorner
				} else {
					char = tcell.RuneVLine
				}
			} else if column == width+x {
				if row == y {
					char = tcell.RuneURCorner
				} else if row == height-1 {
					char = tcell.RuneLRCorner
				} else {
					char = tcell.RuneVLine
				}
			} else if row == y || row == height-1 {
				char = tcell.RuneHLine
			}
			if !r.Hollow || char != 0 {
				point := visualizer.Point{
					X:     x,
					Y:     y,
					Char:  char,
					Style: r.Style,
				}
				points = append(points, point)
			}

		}
	}
	return points
}

//Scale implements the Shape.Scale interface requirement
func (r *Rect) Scale(factor *visualizer.ScaleFactor) {
	if factor == nil {
		return
	}

	origRect := *r
	_, _, newColumns, newRows := r.getRect()
	_, _, origColumns, origRows := r.getRect()

	newColumns = origColumns * factor.ColumnScaleFactor

	newRows = origRows * factor.RowScaleFactor

	r.set(2, newColumns)
	r.set(3, newRows)

	// r[2] = newColumns
	// r[3] = newRows

	log.Debugf("Scaled r by %v from %s to %s",
		factor,
		&origRect,
		r,
	)
}

//GetScaleFactor returns a ScaleFactor that can be used to scale other TermRects
func (r *Rect) GetScaleFactor() *visualizer.ScaleFactor {
	_, _, newColumns, newRows := r.getRect()

	columnsIn := math.ColumnsToInches(newColumns)
	rowsIn := math.RowsToInches(newRows)

	preAspectRatio := columnsIn / rowsIn
	if columnsIn > rowsIn {
		//max out columns, then bring rows to match
		newColumns = math.GetMaxColumns()
		newRows = newColumns / preAspectRatio
	} else {
		//max out rows, then bring columns to match
		newRows = math.GetMaxRows()
		newColumns = preAspectRatio * newRows
	}

	_, _, oldColumns, oldRows := r.getRect()

	columnScaleFactor := newColumns / oldColumns
	rowScaleFactor := newRows / oldRows

	return &visualizer.ScaleFactor{
		ColumnScaleFactor: columnScaleFactor,
		RowScaleFactor:    rowScaleFactor,
	}
}
