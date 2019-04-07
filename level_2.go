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
█    ☲☵      ☵☲☵    █
█ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼ ◼◼☲█
█☲☵      ☲☵   ☲☵  ☲☵█
█ ◼◼☵◼◼ ◼◼ ◼◼ ◼◼ ◼◼☵█
█           ☲☵ exit ↓
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
`

func NewLevel2(g *Game) *Level {

	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)
	tileData := TileMapCellDataMap{
		'b': TileMapCellData{
			ch:                 '💣',
			fgColor:            termbox.ColorGreen,
			bgColor:            termbox.ColorBlack,
			collidesPhysically: true,
			collisionCallback:  nil,
			initCallback: func(selfEntity *Entity) {
				bombOptions := WordOptions{InitCallback: nil, Fg: typedCharacterFg, Bg: typedCharacterBg, CollidesPhysically: true}
				bomb := NewWord(g.Stage, selfEntity.GetPositionX(), selfEntity.GetPositionY(), string('💣'), bombOptions)
				g.Stage.AddTypedEntity(bomb)

				go func() {
					<-time.After(1 * time.Second)
					posX := selfEntity.Position.x
					posY := selfEntity.Position.y
					positions := [][2]int{
						{posX, posY},
						{posX + 1, posY},
						{posX, posY + 1},
						{posX - 1, posY},
						{posX, posY - 1},
					}

					var positionsToBeCleared [][2]int

					for _, pos := range positions {
						if !g.Stage.Canvas.IsInsideOfBoundaries(pos[0], pos[1]) {
							return
						}

						// deliberately using reverse order in two dimensional array :/
						if !ContainsRune([]rune{'◼', '▅', '█'}, g.Stage.LevelInstance.TileMap[pos[1]][pos[0]].Ch) {
							positionsToBeCleared = append(positionsToBeCleared, [2]int{pos[0], pos[1]})
						}
					}

					// clear character and collision
					g.Stage.ClearTileMapCellsAt(positionsToBeCleared)
				}()
			},
		},
		'↓': TileMapCellData{
			ch:                 '↓',
			fgColor:            termbox.ColorGreen,
			bgColor:            termbox.ColorBlack,
			collidesPhysically: false,
			collisionCallback: func() {
				levelInstance := NewLevel3(g)
				g.Stage.SetLevel(levelInstance)
			},
		},
	}

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: level2TileMapString,
		TileData:      tileData,
		InputRunes:    []rune{'b'},
		BlockedKeys:   []termbox.Key{termbox.KeyBackspace},
		Init: func() {
			// load info
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 2 - Bomberman - Vim Modes", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "i: Insert Mode, esc: Back to Normal Mode", explanationOptions)

			hintOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			hint := NewWord(g.Stage, levelHintCoordX, levelHintCoordY, "Type b in Insert Mode to drop a bomb!", hintOptions)

			g.Stage.AddScreenEntity(title, explanation, hint)
		},
	}
}
