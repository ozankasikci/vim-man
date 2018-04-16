package fantasia

import tb "github.com/nsf/termbox-go"

type Canvas [][]tb.Cell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, height)
	for i := range canvas {
		canvas[i] = make([]tb.Cell, width)
	}
	return canvas
}
