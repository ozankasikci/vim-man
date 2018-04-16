package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Entity struct {
	Position point
	Width    int
	Height   int
	Rune     rune
	Cell     *TileMapCell
	Cells    []*TileMapCell
	Stage    *Stage
}

func NewEntity(x, y, w, h int, r rune, fg termbox.Attribute, bg termbox.Attribute, cells []*TileMapCell) *Entity {
	p := point{x, y}
	cell := &TileMapCell{&termbox.Cell{r, fg, bg}, false}
	return &Entity{p, w, h, r, cell, cells, nil}
}

func (e *Entity) SetStage(s *Stage) {
	e.Stage = s
}

func (e *Entity) SetCells(s *Stage) {
	newPositionY := e.Position.y

	for i := 0; i < e.Height; i++ {
		newPositionX := e.Position.x
		if i != 0 {
			newPositionY += 1
		}

		for j := 0; j < e.Width; j++ {
			if j != 0 {
				newPositionX += 1
			}

			if e.Cells != nil {
				index := j

				if e.Width > 0 {
					index = i
				}

				tileMapCell := e.Cells[index]
				s.SetCell(newPositionX, newPositionY, tileMapCell)
			} else {
				tileMapCell := e.Cell
				s.SetCell(newPositionX, newPositionY, tileMapCell)
			}
		}
	}
}

func (e *Entity) Update(s *Stage, event termbox.Event, time time.Time) {
}

func (e *Entity) setPosition(x, y int) {
	e.setPositionX(x)
	e.setPositionY(y)
}

func (e *Entity) setPositionX(x int) {
	e.Position.x = x
}

func (e *Entity) setPositionY(y int) {
	e.Position.y = y
}

func (e *Entity) checkCollision(x, y int) {
	e.Position.y = y
}

func (e *Entity) getPositionX() int {
	return e.Position.x
}

func (e *Entity) getPositionY() int {
	return e.Position.y
}
