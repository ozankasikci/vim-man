package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

const LevelBombermanTileMapString = `
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

func NewLevelBomberman(g *Game) *Level {

	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)
	tileData := TileMapCellDataMap{
		'b': TileMapCellData{
			Ch:                 '💣',
			FgColor:            termbox.ColorGreen,
			BgColor:            termbox.ColorBlack,
			CollidesPhysically: true,
			CollisionCallback:  nil,
			InitCallback: func(selfEntity *Entity) {
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
			Ch:                 '↓',
			FgColor:            termbox.ColorGreen,
			BgColor:            termbox.ColorBlack,
			CollidesPhysically: false,
			CollisionCallback: func() {
				levelInstance := NewLevelExitingVim(g)
				g.Stage.SetLevel(levelInstance)
			},
		},
	}

	level := &Level{
		Game:                 g,
		Entities:             entities,
		TileMapString:        LevelBombermanTileMapString,
		TileData:             tileData,
		InputRunes:           []rune{'b'},
		BlockedKeys:          []termbox.Key{termbox.KeyBackspace, termbox.KeyDelete},
		VimMode:              normalMode,
		TextShiftingDisabled: true,
		Init: func() {
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 4 - VIMBERMAN", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "i: Insert Mode, esc: Back to Normal Mode", explanationOptions)

			hintOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			hint := NewWord(g.Stage, levelHintCoordX, levelHintCoordY, "Type b in Insert Mode to drop a bomb!", hintOptions)

			g.Stage.AddScreenEntity(title, explanation, hint)
		},
	}

	level.InitDefaults()
	return level
}
