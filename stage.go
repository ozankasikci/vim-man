package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

// holds the current level meta data, canvas and entities
type Stage struct {
	Game      *Game
	Level     int
	Fps       float64
	Entities  []Renderer
	BgCell    *termbox.Cell
	Canvas    Canvas
	Width     int
	Height    int
	pixelMode bool
	offsetx   int
	offsety   int
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

func (s *Stage) render() {
	s.setBackgroundCells()
	for _, e := range s.Entities {
		e.SetCells(s)
	}

	termboxSetCells(&s.Canvas)
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
	user := NewUser()
	word := NewWord(20, 30, "Test")
	word.SetStage(s)
	s.Canvas = NewCanvas(10, 10)
	s.AddEntity(user)
	s.AddEntity(word)
}

func (s *Stage) resize(w, h int) {
	s.Width = w
	s.Height = h

	if s.pixelMode {
		s.Height *= 2
	}
	c := NewCanvas(s.Width, s.Height)

	// Copy old data that fits
	for i := 0; i < min(s.Width, len(s.Canvas)); i++ {
		for j := 0; j < min(s.Height, len(s.Canvas[0])); j++ {
			c[i][j] = s.Canvas[i][j]
		}
	}
	s.Canvas = c
}

func (s *Stage) setBackgroundCells() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			s.Canvas[i][j] = *s.BgCell
		}
	}
}

func (s *Stage) SetCell(x, y int, c termbox.Cell) {
	if x >= 0 && x < len(s.Canvas) &&
		y >= 0 && y < len(s.Canvas[0]) {
		s.Canvas[x][y] = c
	}
}

func termboxSetCells(canvas *Canvas) {
	for i, col := range *canvas {
		for j, cell := range col {
			termbox.SetCell(i, j, cell.Ch,
				termbox.Attribute(cell.Fg),
				termbox.Attribute(cell.Bg))
		}
	}

}
