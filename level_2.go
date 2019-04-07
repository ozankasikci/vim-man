package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

const level2TileMapString = `
▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅
█      ☵☲     ☵☲    █
█☲◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█   ☲☵☲☵            █
█ ◼◼☲◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█    ☲☵             █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█☲☵      ☲☵         █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ █
█              exit ↓
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

				go func() {
					<-time.After(1 * time.Second)
					GetLogger().LogValue(selfEntity.Position)
					characterOptions := WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg, CollidesPhysically: false}
					emptyChar1 := NewEmptyCharacter(g.Stage, selfEntity.Position.x, selfEntity.Position.y, characterOptions)
					emptyChar2 := NewEmptyCharacter(g.Stage, selfEntity.Position.x + 1, selfEntity.Position.y, characterOptions)
					g.Stage.AddTypedEntity(emptyChar1, emptyChar2, )
				}()
			},
		},
		'↓': TileMapCellData{
			ch:                 '↓',
			fgColor:            termbox.ColorGreen,
			bgColor:            termbox.ColorBlack,
			collidesPhysically: false,
			collisionCallback: func() {
				levelInstance := NewLevel2(g)
				g.Stage.SetLevel(levelInstance)
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
