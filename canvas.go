package fantasia

import tb "github.com/nsf/termbox-go"

type Canvas [][]tb.Cell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]tb.Cell, height)
	}
	return canvas
}
