package fantasia

import (
	"github.com/nsf/termbox-go"
	"reflect"
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

	tags := []Tag{{"Cursor"}}
	entityOptions := EntityOptions{9999, tags, nil}
	e := NewEntity(s, x, y, 1, 1, ' ', termbox.ColorBlue, termbox.ColorWhite, cells, false, entityOptions)
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
			u.SetPositionY(nextY)
		}
	case 'j':
		nextY := u.GetPositionY() + 1
		if !s.CheckCollision(u.GetPositionX(), nextY) {
			u.SetPositionY(nextY)
		}
	case 'l':
		nextX := u.GetPositionX() + 1
		if !s.CheckCollision(nextX, u.GetPositionY()) {
			u.SetPositionX(nextX)
		}
	case 'h':
		nextX := u.GetPositionX() - 1
		if !s.CheckCollision(nextX, u.GetPositionY()) {
			u.SetPositionX(nextX)
		}
	case 'i':
		if s.LevelInstance.VimMode != insertMode && !s.LevelInstance.InputBlocked {
			s.LevelInstance.VimMode = insertMode
		}
	}
}

func (u *User) handleInsertModeEvents(s *Stage, event termbox.Event) {
	// return on empty event
	if reflect.DeepEqual(event, termbox.Event{}) {
		return
	}

	switch event.Key {
	// switch to normal mode on esc key event
	case termbox.KeyEsc:
		s.LevelInstance.VimMode = normalMode
		return
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		if ContainsTermboxKey(s.LevelInstance.BlockedKeys, termbox.KeyBackspace) ||
			ContainsTermboxKey(s.LevelInstance.BlockedKeys, termbox.KeyBackspace2) {
			return
		}

		characterOptions := WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg}
		character := NewEmptyCharacter(s, u.GetPositionX()-1, u.GetPositionY(), characterOptions)
		if !character.IsInsideOfCanvasBoundaries() {
			return
		}

		s.AddTypedEntity(character)
		u.SetPositionX(u.GetPositionX() - 1)
	default:

		// check if rune input is allowed
		if len(s.LevelInstance.InputRunes) > 0 {
			if !ContainsRune(s.LevelInstance.InputRunes, event.Ch) {
				return
			}
		}

		characterOptions := WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg}
		character := NewWord(s, u.GetPositionX(), u.GetPositionY(), string(event.Ch), characterOptions)
		if !character.IsInsideOfCanvasBoundaries() {
			return
		}

		// type a character and add as typed entity
		s.AddTypedEntity(character)
		u.SetPositionX(u.GetPositionX() + 1)

		if character.InitCallback != nil {
			character.InitCallback()
		}

		if s.LevelInstance.TileData[event.Ch].initCallback != nil {
			s.LevelInstance.TileData[event.Ch].initCallback(character.Entity)
		}
	}
}

func (u *User) Update(s *Stage, event termbox.Event, delta time.Duration) {
	if s.LevelInstance.VimMode == normalMode {
		u.handleNormalModeEvents(s, event)
	} else if s.LevelInstance.VimMode == insertMode {
		u.handleInsertModeEvents(s, event)
	}
}
