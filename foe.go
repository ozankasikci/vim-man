package fantasia

import "github.com/nsf/termbox-go"

const (
	Murloc = 0
)

type Foe struct {
	*Entity
	Type int
}

func NewFoe(x, y int) (u *Foe) {
	e := NewEntity(x, y, 1, 1, '•', termbox.ColorMagenta, termbox.ColorDefault, nil)
	u = &Foe{
		Entity: e,
		Type:   Murloc,
	}
	return
}
