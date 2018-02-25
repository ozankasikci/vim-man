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
	Name     string
	Class    int
	Position point
}

func NewUser() (u *User) {
	e := NewEntity(10, 10, 10, 10, '-', termbox.ColorYellow, termbox.ColorBlue)
	u = &User{
		Entity:   e,
		Name:     "Test",
		Class:    Orc,
		Position: point{20, 20},
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
