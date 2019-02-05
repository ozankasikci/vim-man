package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

// holds the current level meta data, canvas and entities
type Stage struct {
	Game           *Game
	Level          int
	LevelInstance  *Level
	Fps            float64
	CanvasEntities []Renderer // entities to be rendered in canvas
	ScreenEntities []Renderer // entities to be rendered outside of the canvas
	TypedEntities  []Renderer
	BgCell         *termbox.Cell
	Canvas         Canvas
	Width          int
	Height         int
	pixelMode      bool
	offsetx        int
	offsety        int
}

func NewStage(g *Game, level int, fps float64, bgCell *termbox.Cell) *Stage {
	return &Stage{
		Game:           g,
		Level:          level,
		Fps:            fps,
		CanvasEntities: nil,
		ScreenEntities: nil,
		TypedEntities:  nil,
		BgCell:         bgCell,
		Canvas:         nil,
		Width:          0,
		Height:         0,
		pixelMode:      false,
		offsetx:        0,
		offsety:        0,
	}
}

func (s *Stage) AddCanvasEntity(e Renderer) {
	s.CanvasEntities = append(s.CanvasEntities, e)
}

func (s *Stage) AddScreenEntity(e ...Renderer) {
	s.ScreenEntities = append(s.ScreenEntities, e...)
}

func (s *Stage) AddTypedEntity(e ...Renderer) {
	s.TypedEntities = append(s.TypedEntities, e...)
}

func (s *Stage) ClearCanvasEntities() {
	s.CanvasEntities = nil
	s.ScreenEntities = nil
}

func (s *Stage) SetGame(game *Game) {
	s.Game = game
}

// this function handles all the rendering
// sets and renders in order:
// tilemap cells: the background cells
// canvas entities: canvas entities, they have update methods that gets called on stage.update method
// screen cells: cells outside of the canvas such as level label and instructions
func (s *Stage) Render() {
	s.SetCanvasBackgroundCells()

	for i, _ := range s.CanvasEntities {
		e := s.CanvasEntities[i]
		// sets the cells for the entity, so that it overwrites the background cell(s)
		e.SetCells(s)
	}

	s.TermboxSetScreenCells()
	s.TermboxSetCanvasCells()
	s.TermboxSetTypedCells()
	termbox.Flush()
}

func (s *Stage) update(ev termbox.Event, delta time.Duration) {
	for _, e := range s.CanvasEntities {
		e.Update(s, ev, delta)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *Stage) Init() {
	s.SetLevel(NewLevel2(s.Game))
}

func (s *Stage) SetLevel(levelInstance *Level) {
	s.Reset()
	s.LevelInstance = levelInstance
	s.LevelInstance.Init()
	s.LevelInstance.LoadTileMap()
	s.resize(s.LevelInstance.GetTileMapDimensions())

	for _, e := range s.LevelInstance.Entities {
		s.AddCanvasEntity(e)
	}
}

func (s *Stage) Reset() {
	s.ClearCanvasEntities()
	s.Canvas = NewCanvas(10, 10)
}

func (s *Stage) resize(w, h int) {
	s.Width = w
	s.Height = h

	if s.pixelMode {
		s.Height *= 2
	}

	if s.LevelInstance.Height != 0 && s.LevelInstance.Width != 0 {
		s.Width = s.LevelInstance.Width
		s.Height = s.LevelInstance.Height
	}

	c := NewCanvas(s.Width, s.Height)

	// Copy old data that fits
	for i := 0; i < min(s.Height, len(s.Canvas)); i++ {
		for j := 0; j < min(s.Width, len(s.Canvas)); j++ {
			c[i][j] = s.Canvas[i][j]
		}
	}
	s.Canvas = c
}

func (s *Stage) GetDefaultBgCell()  *TermBoxCell{
	return &TermBoxCell{s.BgCell, false, TileMapCellData{}}
}

// sets the background cells to be rendered, this gets rendered first in the render method
// so that other cells can be overwritten into the same location
func (s *Stage) SetCanvasBackgroundCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			if s.LevelInstance.TileMap[i][j].Cell != nil {
				// insert tile map cell
				s.Canvas[i][j] = s.LevelInstance.TileMap[i][j]
			} else {
				//insert default bg cell
				s.Canvas[i][j] = s.GetDefaultBgCell()
			}
		}
	}
}

// calls termbox.setCell, sets the coordinates and the cell attributes
// this does the actual rendering of the characters, thanks to termbox-go <3
func (s *Stage) TermboxSetCell(x, y int, cell *TermBoxCell, offset bool) {
	if offset {
		offsetX, offsetY := s.LevelInstance.GetScreenOffset()
		x += offsetX
		y += offsetY
	}

	termbox.SetCell(x, y, cell.Ch,
		termbox.Attribute(cell.Fg),
		termbox.Attribute(cell.Bg))
}

// sets the cells inside the canvas, offset is being applied in order to keep the canvas in center
func (s *Stage) TermboxSetCanvasCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			cell := row[j]
			// intentionally use j,i in reverse order
			s.TermboxSetCell(j, i, cell, true)
		}
	}
}

// sets the cells outside of the canvas, no offset is being applied
func (s *Stage) TermboxSetScreenCells() {
	for _, e := range s.ScreenEntities {
		for j, _ := range e.GetCells() {
			cell := e.GetCells()[j]
			// intentionally use j,i in reverse order
			offsetX, _ := e.GetScreenOffset()
			x := e.GetPositionX() + j + offsetX
			s.TermboxSetCell(x, e.GetPositionY(), cell, false)
		}
	}
}

func (s *Stage) TermboxSetTypedCells() {
	for _, e := range s.TypedEntities {
		for j, _ := range e.GetCells() {
			cell := e.GetCells()[j]
			// intentionally use j,i in reverse order
			s.TermboxSetCell(e.GetPositionX(), e.GetPositionY(), cell, true)
		}
	}
}

func (s *Stage) CheckCollision(x, y int) bool {
	return s.Canvas.checkCollision(x, y)
}
