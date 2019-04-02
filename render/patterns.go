package render

import (
	"fmt"

	"github.com/gdamore/tcell"
)

var (
	fakeExtraRunes = make([]rune, 0)
)

type shape interface {
	render() []point
}

type point struct {
	x     int
	y     int
	char  rune
	style tcell.Style
}

func (p *point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func draw(screen tcell.Screen, coords []point) {
	for _, coord := range coords {
		screen.SetContent(coord.x, coord.y, coord.char, fakeExtraRunes, coord.style)
	}
}

type Rect struct {
	x      int
	y      int
	width  int
	height int

	hollow bool
	text   *string
	style  tcell.Style
}

func (r *Rect) render() []point {
	var points []point
	return points
}

//RectHollow renders a hollow rectangle
func RectHollow(screen tcell.Screen, x, y, width, height int) {
	rectHollow(screen, x, y, width, height, tcell.StyleDefault)
}

func rectHollow(screen tcell.Screen, x, y, width, height int, style tcell.Style) {
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
			if char != 0 {
				screen.SetContent(column, row, char, fakeExtraRunes, style)
			}
		}
	}
}

func rectHollowWithText(screen tcell.Screen, x, y, width, height int, text string, style tcell.Style) {
	text = fmt.Sprintf("[%s]", text)
	rectHollow(screen, x, y, width, height, style)
	textStart := (width / 2) - (len(text) / 2)
	for i := 0; i < len(text); i++ {
		row := x + i + textStart
		char := rune(text[i])
		screen.SetContent(row, y, char, fakeExtraRunes, style)
	}
}

//RectHollowWithText renders a hollow rectangle with the text centered along the top
func RectHollowWithText(screen tcell.Screen, x, y, width, height int, text string) {
	rectHollowWithText(screen, x, y, width, height, text, tcell.StyleDefault)
}

//RectHollowWithTextAndStyle renders a hollow rectangle with the text centered along the top and the style rules applied to the text
func RectHollowWithTextAndStyle(screen tcell.Screen, x, y, width, height int, text string, style tcell.Style) {
	rectHollowWithText(screen, x, y, width, height, text, style)
	// textStart := (width / 2) - (len(text) / 2)
	// for i := 0; i < len(text); i++ {
	// row := x + i + textStart
	// char := rune(text[i])
	// screen.SetContent(row, y, char, fakeExtraRunes, style)
	// }
}
