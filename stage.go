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
	Entities      []Renderer
	BgCell        *termbox.Cell
	Canvas        Canvas
	Width         int
	Height        int
	pixelMode     bool
	offsetx       int
	offsety       int
}

func NewStage(level int, fps float64, bgCell *termbox.Cell) *Stage {
	return &Stage{
		Game:      nil,
		Level:     level,
		Fps:       fps,
		Entities:  nil,
		BgCell:    bgCell,
		Canvas:    nil,
		Width:     0,
		Height:    0,
		pixelMode: false,
		offsetx:   0,
		offsety:   0,
	}
}

func (s *Stage) AddEntity(e Renderer) {
	s.Entities = append(s.Entities, e)
}

func (s *Stage) SetGame(game *Game) {
	s.Game = game
}

func (s *Stage) Render() {
	s.SetBackgroundCells()
	for _, e := range s.Entities {
		e.SetCells(s)
	}

	s.TermboxSetCells()
	termbox.Flush()
}

func (s *Stage) update(ev termbox.Event, delta time.Duration) {
	for _, e := range s.Entities {
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
	level1 := NewLevel1()
	s.LevelInstance = level1
	s.LevelInstance.LoadTileMap()
	s.resize(s.LevelInstance.GetTileMapDimensions())
	s.AddEntity(level1.Entities[0])
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

func (s *Stage) SetBackgroundCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			if s.LevelInstance.TileMap[i][j].Cell != nil {
				s.Canvas[i][j] = *s.LevelInstance.TileMap[i][j].Cell
			} else {
				s.Canvas[i][j] = *s.BgCell
			}
		}
	}
}

func (s *Stage) SetCell(x, y int, c termbox.Cell) {
	if x >= 0 && x < len(s.Canvas) && y >= 0 && y < len(s.Canvas[0]) {
		s.Canvas[y][x] = c
	}
}

func (s *Stage) TermboxSetCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			cell := row[j]
			termbox.SetCell(j, i, cell.Ch,
				termbox.Attribute(cell.Fg),
				termbox.Attribute(cell.Bg))
		}
	}
}
