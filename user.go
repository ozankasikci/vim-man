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
	cells := []*TermBoxCell{
		{&termbox.Cell{'▒', termbox.ColorGreen, bgColor}, false, TileMapCellData{}},
	}

	e := NewEntity(s, x, y, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells, false)
	u = &User{
		Entity: e,
	}
	return
}

func (u *User) handleNormalModeEvents(s *Stage, event termbox.Event) {
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
	case 'i':
		if s.LevelInstance.VimMode != insertMode {
			s.LevelInstance.VimMode = insertMode
		}
	}
}

func (u *User) handleInsertModeEvents(s *Stage, event termbox.Event) {
	GetLogger().WriteFile("entered handleInsertModeEvents")
	switch event.Key {
	case termbox.KeyEsc:
		GetLogger().WriteFile("esc pressed")
		s.LevelInstance.VimMode = normalMode
		return
	}

	if event.Ch == 'o' {
		//cell := &TermBoxCell{
		//	Cell: &termbox.Cell{'a', termbox.ColorGreen, termbox.ColorGreen},
		//	collidesPhysically: false,
		//	cellData: TileMapCellData{},
		//}
		//s.SetCanvasCell(0, 0, cell)
		GetLogger().WriteFile("setting canvas cell in insert mode")
		GetLogger().WriteFile("o pressed")
	}

}

func (u *User) Update(s *Stage, event termbox.Event, delta time.Duration) {
	if event.Ch == 'E' {
		GetLogger().LogValue(event.Key)
	}
	if s.LevelInstance.VimMode == normalMode {
		GetLogger().WriteFile("if s.LevelInstance.VimMode == normalMode is true")
		u.handleNormalModeEvents(s, event)
	} else if s.LevelInstance.VimMode == insertMode {
		GetLogger().WriteFile("if s.LevelInstance.VimMode == insertMod is true")
		u.handleInsertModeEvents(s, event)
	}

}
