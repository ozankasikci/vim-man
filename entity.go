package fantasia

import "github.com/nsf/termbox-go"

type Entity struct {
	Position point
	Width    int
	Height   int
	Rune     rune
	Cell     termbox.Cell
}

func NewEntity(x, y, w, h int, r rune, fg termbox.Attribute, bg termbox.Attribute) *Entity {
	p := point{x, y}
	cell := termbox.Cell{r, fg, bg}
	return &Entity{p, w, h, r, cell}
}

func (e *Entity) Render(s *Stage) {
	renderSquare(e.Position.x, e.Position.y, e.Width, e.Height, e.Cell)
}

func (e *Entity) Update(s *Stage, event termbox.Event) {
}
