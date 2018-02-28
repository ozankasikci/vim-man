package fantasia

import "github.com/nsf/termbox-go"

type Current struct {
	*Entity
}

func NewCurrent(x, y int) (c *Current) {
	e := NewEntity(x, y, 1, 1, '~', termbox.ColorBlue, termbox.ColorBlack)
	c = &Current{
		Entity: e,
	}
	return
}

func (c *Current) Update(s *Stage, event termbox.Event)  {
	c.Entity.Position.x = c.Entity.Position.x + 1
}