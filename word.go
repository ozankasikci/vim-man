package fantasia

import (
	"github.com/nsf/termbox-go"
)

type Word struct {
	*Entity
	Content string
}

func ConvertStringToCells(s string) []termbox.Cell{
	var arr []termbox.Cell
	//for _, rune := range s {
	//	arr = append(arr, termbox.Cell{rune, termbox.ColorGreen, termbox.ColorWhite})
	//}
	for i := 0; i < len([]rune(s)); i++  {
		cell := termbox.Cell{[]rune(s)[i], termbox.ColorGreen, termbox.ColorWhite}
		arr = append(arr, cell)
	}
	//for i, _ := range s {
	//	cell := termbox.Cell{[]rune(s)[i], termbox.ColorGreen, termbox.ColorWhite}
	//	arr = append(arr, cell)
	//}
	return arr
}

func NewWord(x, y int, content string) (w *Word) {
	cells := ConvertStringToCells(content)
	e := NewEntity(x, y, 4, 1, ' ', termbox.ColorMagenta, bgColor, cells)
	w = &Word{
		Entity: e,
		Content: content,
	}
	return
}
