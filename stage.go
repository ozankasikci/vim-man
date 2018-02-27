package fantasia

import (
	"github.com/nsf/termbox-go"
)

type Stage struct {
	Level     int
	Fps       float64
	Entities  []Renderer
	Canvas    Canvas
	width     int
	height    int
	pixelMode bool
	offsetx   int
	offsety   int
}

func NewStage(level int, fps float64) *Stage {
	return &Stage{
		level,
		fps,
		nil,
		nil,
		0,
		0,
		false,
		0,
		0,
	}
}

func (s *Stage) AddEntity(e Renderer) {
	s.Entities = append(s.Entities, e)
}

func (s *Stage) render() {
	s.RenderBackground()
	for _, e := range s.Entities {
		e.Render(s)
	}
	termbox.Flush()
}

func (s *Stage) update(ev termbox.Event) {
	for _, e := range s.Entities {
		e.Update(s, ev)
	}
}

func renderSquare(x, y, w, h int, cell termbox.Cell) {
	for iy := 0; iy < h; iy++ {
		for ix := 0; ix < w; ix++ {
			termbox.SetCell(x+ix, y+iy, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func renderWord(w word, active bool) {
	runes := []rune(w.text)
	for i, r := range runes {
		fgColor := termbox.ColorDefault
		if i == w.cursor && active {
			fgColor = termbox.ColorRed
		}
		termbox.SetCell(w.location.x+i, w.location.y, r, fgColor, termbox.ColorDefault)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (s *Stage) Init() {
	user := NewUser()
	s.Canvas = NewCanvas(10, 10)
	s.AddEntity(user)
}

func (s *Stage) resize(w, h int) {
	s.width = w
	s.height = h

	if s.pixelMode {
		s.height *= 2
	}
	c := NewCanvas(s.width, s.height)

	// Copy old data that fits
	for i := 0; i < min(s.width, len(s.Canvas)); i++ {
		for j := 0; j < min(s.height, len(s.Canvas[0])); j++ {
			c[i][j] = s.Canvas[i][j]
		}
	}
	s.Canvas = c
}

func (s *Stage) RenderBackground() {
	for i, row := range s.Canvas {
		for j, _ := range row {
			s.Canvas[i][j] = termbox.Cell{
				Ch: '~',
				Fg: termbox.ColorGreen,
				Bg: termbox.ColorBlack,
			}
		}
	}
}
