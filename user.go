package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Class int

type User struct {
	*Entity
	Name string
}

func NewUser() (u *User) {
	cells := []*termbox.Cell{
		&termbox.Cell{'▓', termbox.ColorGreen, bgColor},
	}

	e := NewEntity(0, 0, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells)
	u = &User{
		Entity: e,
		Name:   "Test",
	}
	return
}

func (u *User) Update(s *Stage, event termbox.Event, delta time.Duration) {
	switch event.Ch {
	case 'k':
		u.setPositionY(u.Entity.Position.y - 1)
	case 'j':
		u.setPositionY(u.Entity.Position.y + 1)
	case 'l':
		u.setPositionX(u.Entity.Position.x + 1)
	case 'h':
		u.setPositionX(u.Entity.Position.x - 1)
	}
}
