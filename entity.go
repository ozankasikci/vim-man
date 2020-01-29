package vimman

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Direction int

const (
	horizontal Direction = iota
	vertical
)

type Entity struct {
	Stage        *Stage
	Position     Point
	Width        int
	Height       int
	Rune         rune
	Cell         *TermBoxCell
	Cells        []*TermBoxCell
	DrawPriority int
	Tags         []Tag
	InitCallback func()
}

type EntityOptions struct {
	DrawPriority int
	Tags         []Tag
	InitCallback func()
}

type Tag struct {
	Name string
}

func NewEntity(s *Stage, x, y, w, h int, r rune, fg termbox.Attribute, bg termbox.Attribute, cells []*TermBoxCell, collidesPhysically bool, options EntityOptions) *Entity {
	drawPriority, tags, initCallback := options.DrawPriority, options.Tags, options.InitCallback
	p := Point{x, y}
	cell := &TermBoxCell{&termbox.Cell{r, fg, bg}, collidesPhysically, TileMapCellData{}}
	return &Entity{s, p, w, h, r, cell, cells, drawPriority, tags, initCallback}
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
					s.Canvas.OverWriteCanvasCell(newPositionX, newPositionY, tileMapCell)
				}
			} else {
				tileMapCell := e.Cell
				s.Canvas.OverWriteCanvasCell(newPositionX, newPositionY, tileMapCell)
			}
		}
	}
}

func (e *Entity) GetCells() []*TermBoxCell {
	return e.Cells
}

func (e *Entity) Update(s *Stage, event termbox.Event, time time.Duration) {
}

func (e *Entity) SetPosition(x, y int) {
	e.SetPositionX(x)
	e.SetPositionY(y)
}

func (e *Entity) SetPositionX(x int) {
	e.Position.x = x
}

func (e *Entity) SetPositionY(y int) {
	e.Position.y = y
}

func (e *Entity) CheckCollision(x, y int) {
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

func (e *Entity) Destroy() {
}

func (e *Entity) GetDrawPriority() int {
	return e.DrawPriority
}

func (e *Entity) GetTags() []Tag {
	return e.Tags
}

func (e *Entity) IsInsideOfCanvasBoundaries() bool {
	return e.GetStage().Canvas.IsInsideOfBoundaries(e.GetPositionX(), e.GetPositionY())
}
