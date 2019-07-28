package fantasia

import "github.com/nsf/termbox-go"

const LevelBasicMovementTileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
  Try to find the exit  |        |        |     |     |     |
+  +--+--+--+--+--+  +--+  +--+  +  +  +  +  +  +  +  +--+  +
|           |     |     |     |  |  |  |     |  |  |  |     |
+--+--+  +--+  +  +--+  +--+  +  +--+  +  +--+  +  +  +  +  +
|        |     |     |     |  |     |  |     |     |     |  |
+  +--+--+  +  +--+--+--+  +--+--+  +  +--+--+--+--+--+--+  +
|  |     |  |  |           |        |           |     |     |
+  +  +  +  +--+  +--+--+--+  +--+--+--+--+--+  +  +  +  +--+
|     |  |  |     |     |     |     |        |     |  |     |
+--+--+  +  +  +--+  +  +  +  +--+  +  +--+  +--+--+  +--+  +
|     |  |     |     |     |        |  |  |  |     |  |     |
+  +  +  +--+  +  +--+  +--+  +--+--+  +  +  +  +--+  +  +--+
|  |  |     |  |  |     |  |  |        |  |        |  |     |
+  +  +--+  +  +  +--+--+  +  +  +--+--+  +--+--+  +  +--+--+
|  |        |  |     |     |           |  |     |  |        |
+  +--+--+--+--+--+  +  +  +--+--+--+  +  +  +  +  +--+--+  +
|  |                 |  |  |     |     |     |  |           |
+  +--+  +--+  +--+--+  +  +  +  +  +--+--+  +  +--+--+--+--+
|        |     |        |     |              |         exit ↓ 
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
`

func NewLevelBasicMovement(g *Game) *Level {
	// create user
	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)

	tileData := TileMapCellDataMap{
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
		Game:          g,
		Entities:      entities,
		TileMapString: LevelBasicMovementTileMapString,
		TileData:      tileData,
		InputBlocked:  true,
		VimMode:       normalMode,
		Init: func() {
			// load info
			titleOptions := WordOptions{
				InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 1 - MOVING THE CURSOR", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "J: down, H: left, K: up, L: right", explanationOptions)

			g.Stage.AddScreenEntity(title, explanation)
		},
	}

	level.InitDefaults()
	return level
}
