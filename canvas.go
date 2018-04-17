package fantasia

//import tb "github.com/nsf/termbox-go"

type Canvas [][]*TileMapCell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, height)
	for i := range canvas {
		canvas[i] = make([]*TileMapCell, width)
	}
	return canvas
}

func (c Canvas) getCellAt(x, y int) *TileMapCell {
	return c[y][x]
}

func (c Canvas) checkCollision(x, y int) bool {
	if x < 0 || y < 0 || y >= len(c) || x >= len(c[0]) {
		return true
	}

	if c[y][x] == nil {
		return true
	}

	return c[y][x].collides
}
