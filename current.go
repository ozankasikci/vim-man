package fantasia

import "github.com/nsf/termbox-go"

type Current struct {
	*Entity
}

func NewCurrent(x, y int) (c *Current) {
	e := NewEntity(x, y, 2, 3, '~', termbox.ColorBlue, termbox.ColorBlack, nil)
	c = &Current{
		Entity: e,
	}
	return
}

func (c *Current) Update(s *Stage, event termbox.Event) {
	c.Entity.Position.x = c.Entity.Position.x + 1
}
