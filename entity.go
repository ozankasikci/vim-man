package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Entity struct {
	Stage    *Stage
	Position point
	Width    int
	Height   int
	Rune     rune
	Cell     *TileMapCell
	Cells    []*TileMapCell
}

func NewEntity(s *Stage, x, y, w, h int, r rune, fg termbox.Attribute, bg termbox.Attribute, cells []*TileMapCell) *Entity {
	p := point{x, y}
	cell := &TileMapCell{&termbox.Cell{r, fg, bg}, false}
	return &Entity{s, p, w, h, r, cell, cells}
}

func (e *Entity) SetStage(s *Stage) {
	e.Stage = s
}

func (e *Entity) GetStage() *Stage {
	return e.Stage
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

				tileMapCell := e.Cells[index]
				if len(e.Cells) > index {
					s.SetCanvasCell(newPositionX, newPositionY, tileMapCell)
				}
			} else {
				tileMapCell := e.Cell
				s.SetCanvasCell(newPositionX, newPositionY, tileMapCell)
			}
		}
	}
}

func (e *Entity) GetCells() []*TileMapCell {
	return e.Cells
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

func (e *Entity) GetPositionX() int {
	return e.Position.x
}

func (e *Entity) GetPositionY() int {
	return e.Position.y
}

func (e *Entity) GetPosition() (int, int) {
	return e.Position.x, e.Position.y
}

func (e *Entity) GetScreenOffset() (int, int) {
	screenWidth, screenHeight := e.Stage.Game.getScreenSize()
	return (screenWidth - e.Width) / 2, (screenHeight - e.Height) / 2
}
