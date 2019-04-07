package fantasia

import "os"

//import tb "github.com/nsf/termbox-go"

type Canvas [][]*TermBoxCell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, height)
	for i := range canvas {
		canvas[i] = make([]*TermBoxCell, width)
	}
	return canvas
}

func (c Canvas) GetCellAt(x, y int) *TermBoxCell {
	return c[y][x]
}

func (c Canvas) CheckCollision(x, y int) bool {
	// check if out of boundaries
	if x < 0 || y < 0 || y >= len(c) || x >= len(c[0]) {
		return true
	}

	if c[y][x] == nil {
		return true
	}

	c[y][x].cellData.collisionCallback()

	if os.Getenv("DEBUG") == "1" {
		return false
	}

	return c[y][x].collidesPhysically
}

func (c *Canvas) OverWriteCanvasCell(x, y int, termboxCell *TermBoxCell) {
	if x >= 0 && x < len((*c)[0]) && y >= 0 && y < len((*c)) {
		// intentionally use x,y in reverse order
		(*c)[y][x] = termboxCell
	}
}

func (c *Canvas) SetCellAt(row, column int, cell *TermBoxCell) {
	(*c)[row][column] = cell
}

func (c *Canvas) IsInsideOfBoundaries(x, y int) bool {
	return x >= 0 && x < len((*c)[0]) && y >= 0 && y < len(*c)
}