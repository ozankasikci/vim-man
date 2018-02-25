package fantasia

import "github.com/nsf/termbox-go"

type Stage struct {
	Level    int
	Fps      float64
	Entities []Renderer
	Canvas   Canvas
}

func NewStage(level int, fps float64) *Stage {
	return &Stage{level, fps, nil, nil}
}

func (s *Stage) AddEntity(e Renderer) {
	s.Entities = append(s.Entities, e)
}

func (s *Stage) render() {
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
