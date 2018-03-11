package fantasia

import "github.com/nsf/termbox-go"

type Entity struct {
	Position point
	Width    int
	Height   int
	Rune     rune
	Cell     *termbox.Cell
	Cells    []termbox.Cell
	Stage    *Stage
}

func NewEntity(x, y, w, h int, r rune, fg termbox.Attribute, bg termbox.Attribute, cells []termbox.Cell) *Entity {
	p := point{x, y}
	cell := &termbox.Cell{r, fg, bg}
	return &Entity{p, w, h, r, cell, cells, nil}
}

func (e *Entity) SetStage(s *Stage)  {
	e.Stage = s
}

func (e *Entity) SetCells(s *Stage) {
	newPositionX := e.Position.x

	for i := 0; i < e.Width; i++ {
		newPositionY := e.Position.y
		newPositionX += 1

		for j := 0; j < e.Height; j++ {
			newPositionY += 1

			if e.Cells != nil {
				index := j

				if e.Width > 0 {
					index = i
				}

				s.SetCell(newPositionX, newPositionY, e.Cells[index])
			} else {
				s.SetCell(newPositionX, newPositionY, *e.Cell)
			}
		}
	}
}

func (e *Entity) Update(s *Stage, event termbox.Event) {
}
