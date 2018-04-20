package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Class int

type User struct {
	*Entity
}

func NewUser(s *Stage, x, y int) (u *User) {
	cells := []*TileMapCell{
		{&termbox.Cell{'▓', termbox.ColorGreen, bgColor}, false},
	}

	e := NewEntity(s, x, y, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells, false)
	u = &User{
		Entity: e,
	}
	return
}

func (u *User) Update(s *Stage, event termbox.Event, delta time.Duration) {
	switch event.Ch {
	case 'k':
		nextY := u.GetPositionY() - 1
		if !s.CheckCollision(u.GetPositionX(), nextY) {
			u.setPositionY(nextY)
		}
	case 'j':
		nextY := u.GetPositionY() + 1
		if !s.CheckCollision(u.GetPositionX(), nextY) {
			u.setPositionY(nextY)
		}
	case 'l':
		nextX := u.GetPositionX() + 1
		if !s.CheckCollision(nextX, u.GetPositionY()) {
			u.setPositionX(nextX)
		}
	case 'h':
		nextX := u.GetPositionX() - 1
		if !s.CheckCollision(nextX, u.GetPositionY()) {
			u.setPositionX(nextX)
		}
	}
}
