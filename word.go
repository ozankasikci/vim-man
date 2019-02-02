package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Word struct {
	*Entity
	Content   string
	Speed     float64
	Direction Direction
}

func ConvertStringToCells(s string) []*TermBoxCell {
	var arr []*TermBoxCell

	for i := 0; i < len([]rune(s)); i++ {
		cell := &TermBoxCell{
			Cell: &termbox.Cell{
				[]rune(s)[i],
				termbox.ColorGreen,
				termbox.ColorBlack,
			},
			collidesPhysically: false,
			cellData: TileMapCellData{}}

		arr = append(arr, cell)
	}

	return arr
}

func NewWord(s *Stage, x, y int, content string) *Word {
	cells := ConvertStringToCells(content)
	e := NewEntity(s, x, y, len(content), 1, ' ', termbox.ColorMagenta, termbox.ColorBlack, cells, false)
	return &Word{
		Entity:    e,
		Content:   content,
		Direction: horizontal,
	}
}

func (w *Word) Update(s *Stage, event termbox.Event, delta time.Duration) {
}
