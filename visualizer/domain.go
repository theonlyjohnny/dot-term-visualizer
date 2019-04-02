package visualizer

import (
	"fmt"

	"github.com/awalterschulze/gographviz/ast"
	"github.com/gdamore/tcell"
)

//An ActionFunc lives within an Element and is used as the directions for how an Element should render
type ActionFunc func(ast.Stmt, *Screen, *GraphStorage)

//A Graph is an Element that groups other Elements together, and is rendered
type Graph struct {
	*Element
	Elements []*Element
}

//A Screen is a composed version of a tcell.Screen, and is generally what content is rendered into
type Screen struct {
	tcell.Screen
	ExitChan chan<- error
}

//An Element is used to draw an ast.Stmt onto a Screen
type Element struct {
	Stmt       ast.Stmt
	ActionFunc ActionFunc
	Storage    *GraphStorage
}

//GraphStorage is used to share information between Elements being rendered as part of the same Graph
type GraphStorage struct {
	ScaleFactor *ScaleFactor
}

//ScaleFactor is used to scale all Elements of the same Graph by the same amount
type ScaleFactor struct {
	ColumnScaleFactor float64
	RowScaleFactor    float64
}

func (sf *ScaleFactor) String() string {
	return fmt.Sprintf("[col:%g * row:%g]",
		sf.ColumnScaleFactor,
		sf.RowScaleFactor,
	)
}

//Draw is called with a tcell.Screen and will cause the associated ast.Stmt to be added to the view however it should be
func (e *Element) Draw(s *Screen) {
	if e.ActionFunc != nil {
		e.ActionFunc(e.Stmt, s, e.Storage)
	}
}

//A Shape is an interface which returns Points to draw
type Shape interface {
	Render() []Point
	Scale(*ScaleFactor)
}

//A Point represents a single character in a single (x,y) point, to be rendered to a Screen
type Point struct {
	X     int
	Y     int
	Char  rune
	Style *tcell.Style
}
