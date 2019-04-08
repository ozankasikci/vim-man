package fantasia

import "github.com/nsf/termbox-go"

//import "github.com/nsf/termbox-go"

const LevelExitingVimTileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
  Try to find the exit  |        |        |     |     |     |
+  +--+--+--+--+--+  +--+  +--+  +  +  +  +  +  +  +  +--+  +
|           |     |     |     |  |  |  |     |  |  |  |     |
+--+--+  +--+  +  +--+  +--+  +  +--+  +  +--+  +  +  +  +  +
|        |     |     |     |  |     |  |     |     |     |  |
+  +--+--+  +  +--+--+--+  +--+--+  +  +--+--+--+--+--+--+  +
|  |     |  |  | if you can|        |           |     |     |
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
+  +--+  +--+  +--+--+  +  +  +  +  +--+--+  +--+--+--+--+--+
|        |     |        |     |              |         exit ↓ 
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
`

func NewLevelExitingVim(g *Game) *Level {
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
			CollisionCallback: func() {},
		},
	}

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: LevelExitingVimTileMapString,
		TileData:      tileData,
		InputBlocked:  true,
		VimMode:       normalMode,
		Init: func() {
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 3 - EXITING VIM", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "Same level? It's a trap. You need learn how to exit vim", explanationOptions)

			hintOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			hint := NewWord(g.Stage, levelHintCoordX, levelHintCoordY, "Type colon ':', then 'q', press enter", hintOptions)

			g.Stage.AddScreenEntity(title, explanation, hint)
		},
	}
}
