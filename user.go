package fantasia

import "github.com/nsf/termbox-go"

type Class int

const (
	Orc   = 0
	Elf   = 1
	Dwarf = 2
	Bear  = 3
)

type User struct {
	*Entity
	Name  string
	Class int
}

func NewUser() (u *User) {
	cells := []termbox.Cell{
		termbox.Cell{'▓', termbox.ColorGreen, bgColor},
		termbox.Cell{'┬', termbox.ColorGreen, bgColor},
	}

	e := NewEntity(10, 10, 1, 2, ' ', termbox.ColorBlue, termbox.ColorWhite, cells)
	u = &User{
		Entity: e,
		Name:   "Test",
		Class:  Orc,
	}
	return
}

func (u *User) Update(s *Stage, event termbox.Event) {
	switch event.Key {
	case termbox.KeyArrowUp:
		u.Entity.Position.y = u.Entity.Position.y - 1
	case termbox.KeyArrowDown:
		u.Entity.Position.y = u.Entity.Position.y + 1
	case termbox.KeyArrowRight:
		u.Entity.Position.x = u.Entity.Position.x + 1
	case termbox.KeyArrowLeft:
		u.Entity.Position.x = u.Entity.Position.x - 1
	}
}
