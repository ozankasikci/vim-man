package fantasia

import "github.com/nsf/termbox-go"

const level2TileMapString = `
▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅
█                   █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█                   █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█                   █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█                   █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█                   █
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
`

func NewLevel2(g *Game) *Level {
	// create user
	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)
	tileData := TileMapCellDataMap{
		'b': TileMapCellData{
			ch:                 '💣',
			fgColor:            termbox.ColorGreen,
			bgColor:            termbox.ColorBlack,
			collidesPhysically: true,
			collisionCallback: nil,
			initCallback: func(selfEntity *Entity) {
				bombOptions := WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg, CollidesPhysically: true}
				bomb := NewWord(g.Stage, selfEntity.GetPositionX(), selfEntity.GetPositionY(), string('💣'), bombOptions)
                g.Stage.AddTypedEntity(bomb)
			},
		},
	}

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: level2TileMapString,
		TileData:      tileData,
		Init: func() {
			// load info
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 2 - Vim Modes", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "i: Insert Mode, esc: Back to Normal Mode", explanationOptions)

			g.Stage.AddScreenEntity(title, explanation)
		},
	}
}
