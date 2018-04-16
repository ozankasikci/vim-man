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
	cells := []*TileMapCell{
		{&termbox.Cell{'▓', termbox.ColorGreen, bgColor}, false},
	}

	e := NewEntity(1, 0, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells)
	u = &User{
		Entity: e,
		Name:   "Test",
	}
	return
}

func (u *User) Update(s *Stage, event termbox.Event, delta time.Duration) {
	switch event.Ch {
	case 'k':
		nextY := u.getPositionY() - 1
		if !s.CheckCollision(u.getPositionX(), nextY) {
			u.setPositionY(nextY)
		}
	case 'j':
		nextY := u.getPositionY() + 1
		if !s.CheckCollision(u.getPositionX(), nextY) {
			u.setPositionY(nextY)
		}
	case 'l':
		nextX := u.getPositionX() + 1
		if !s.CheckCollision(nextX, u.getPositionY()) {
			u.setPositionX(nextX)
		}
	case 'h':
		nextX := u.getPositionX() - 1
		if !s.CheckCollision(nextX, u.getPositionY()) {
			u.setPositionX(nextX)
		}
	}
}
