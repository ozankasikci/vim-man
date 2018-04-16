package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

const (
	horizontal = iota
	vertical
)

type Word struct {
	*Entity
	Content             string
	Speed               float64
	Direction           int
	TimeSinceLastUpdate float64
}

func ConvertStringToCells(s string) []*TileMapCell {
	var arr []*TileMapCell

	for i := 0; i < len([]rune(s)); i++ {
		cell := &TileMapCell{&termbox.Cell{[]rune(s)[i], termbox.ColorGreen, termbox.ColorBlack}, false}
		arr = append(arr, cell)
	}

	return arr
}

func NewWord(x, y int, content string) (w *Word) {
	cells := ConvertStringToCells(content)
	e := NewEntity(x, y, 4, 1, ' ', termbox.ColorMagenta, bgColor, cells)
	w = &Word{
		Entity:              e,
		Content:             content,
		Speed:               0.20,
		Direction:           horizontal,
		TimeSinceLastUpdate: 0,
	}
	return
}

func (w *Word) Update(s *Stage, event termbox.Event, delta time.Duration) {
	w.TimeSinceLastUpdate += delta.Seconds()

	if w.Speed > w.TimeSinceLastUpdate {
		return
	}

	w.Position.x = w.Position.x + 1
	w.TimeSinceLastUpdate = 0
}
