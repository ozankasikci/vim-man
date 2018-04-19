package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

// holds the current level meta data, canvas and entities
type Stage struct {
	Game          *Game
	Level         int
	LevelInstance *Level
	Fps           float64
	CanvasEntities []Renderer
	ScreenEntities []Renderer
	BgCell        *termbox.Cell
	Canvas        Canvas
	Width         int
	Height        int
	pixelMode     bool
	offsetx       int
	offsety       int
}

func NewStage(g *Game, level int, fps float64, bgCell *termbox.Cell) *Stage {
	return &Stage{
		Game:      g,
		Level:     level,
		Fps:       fps,
		CanvasEntities:  nil,
		ScreenEntities:  nil,
		BgCell:    bgCell,
		Canvas:    nil,
		Width:     0,
		Height:    0,
		pixelMode: false,
		offsetx:   0,
		offsety:   0,
	}
}

func (s *Stage) AddCanvasEntity(e Renderer) {
	s.CanvasEntities = append(s.CanvasEntities, e)
}

func (s *Stage) AddScreenEntity(e Renderer) {
	s.ScreenEntities = append(s.ScreenEntities, e)
}

func (s *Stage) SetGame(game *Game) {
	s.Game = game
}

// this function handles all the rendering
// sets and renders of the entity and tilemap cells
func (s *Stage) Render() {
	s.SetCanvasBackgroundCells()
	for i, _ := range s.CanvasEntities {
		e := s.CanvasEntities[i]
		e.SetCells(s)
	}

	s.TermboxSetScreenCells()
	s.TermboxSetCanvasCells()
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
	s.Canvas = NewCanvas(10, 10)
	s.LevelInstance = NewLevel1(s.Game)
	s.LevelInstance.LoadTileMap()
	s.resize(s.LevelInstance.GetTileMapDimensions())

	for _, e := range s.LevelInstance.Entities {
		s.AddCanvasEntity(e)
	}
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

func (s *Stage) SetCanvasBackgroundCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			if s.LevelInstance.TileMap[i][j].Cell != nil {
				// insert tile map cell
				s.Canvas[i][j] = s.LevelInstance.TileMap[i][j]
			} else {
				 //insert default bg cell
				s.Canvas[i][j] = &TileMapCell{s.BgCell, false}
			}
		}
	}
}

func (s *Stage) SetCanvasCell(x, y int, c *TileMapCell) {
	if x >= 0 && x < len(s.Canvas[0]) && y >= 0 && y < len(s.Canvas) {
		// intentionally use x,y in reverse order
		s.Canvas[y][x] = c
	} else {
		lg.LogValue(x >= 0, x < len(s.Canvas))
	}
}

func (s *Stage) TermboxSetCell(x, y int, cell *TileMapCell) {
	termbox.SetCell(x, y, cell.Ch,
		termbox.Attribute(cell.Fg),
		termbox.Attribute(cell.Bg))
}

func (s *Stage) TermboxSetCanvasCells () {
	offsetX, offsetY := s.LevelInstance.GetScreenOffset()

	for i, row := range s.Canvas {
		for j, _ := range row {
			cell := row[j]
			// intentionally use j,i in reverse order
			s.TermboxSetCell(j + offsetX, i + offsetY, cell)
		}
	}
}

func (s *Stage) TermboxSetScreenCells () {
	for _, e := range s.ScreenEntities {
		for j, _ := range e.GetCells() {
			cell := e.GetCells()[j]
			// intentionally use j,i in reverse order
			offsetX, _ := e.GetScreenOffset()
			x := e.GetPositionX() + j + offsetX
			s.TermboxSetCell(x, e.GetPositionY(), cell)
		}
	}
}

func (s *Stage) CheckCollision(x, y int) bool {
	return s.Canvas.checkCollision(x, y)
}
