package fantasia

import "github.com/nsf/termbox-go"

type Level2 struct {
	*Level
	Stage *Stage
}

func NewLevel2() *Level {
	user := NewUser()
	cells := []termbox.Cell{
		termbox.Cell{'▓', termbox.ColorGreen, bgColor},
	}

	e := NewEntity(20, 20, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells)
	u := &User{
		Entity: e,
		Name:   "Test",
	}
	var entities []Renderer
	entities = append(entities, user, u)

	return &Level{
		Entities: entities,
	}
}
