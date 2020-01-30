package vimman

import "github.com/nsf/termbox-go"

const LevelTextEditingTileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|~                                                          |
|~                                                          |
|~                                                          |
|~  1) DELETION - Press "x" to delete the character under   |
|~              the cursor  in normal mode.                 |
|~                                                          |
|~               Yyou shalll noot paass                     |
|~                                                          |
|~  2) INSERTION - Move the cursor after the character      |
|~  where the text should be inserted, press i and type     |
|~                                                          |
|~                   Yu shll nt pss                         |
|~                                                          |
|~  3) APPENDING - Move the curser before the character     |
|~  where the text should be appended, press a and type     |
|~                                                          |
|~                    Yo shal no pas                        |
|~                                                          |
|~                                                          |
|~                                                          |
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
`

func NewLevelTextEditing(g *Game) *Level {
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
			CollisionCallback:  func() {},
		},
	}

	level := &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: LevelTextEditingTileMapString,
		TileData:      tileData,
		VimMode:       normalMode,
		Init: func() {
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 3 - TEXT EDITING", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "Delete, insert, append.", explanationOptions)

			hintOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			hint := NewWord(g.Stage, levelHintCoordX, levelHintCoordY, "Complete the 3 steps below to proceed to the next level.", hintOptions)

			g.Stage.AddScreenEntity(title, explanation, hint)
		},
		//ColonLineCallbacks: map[string]func(*Game) {
		//	"q": func(g *Game) {
		//		levelInstance := NewLevelBomberman(g)
		//		g.Stage.SetLevel(levelInstance)
		//	},
		//},
	}

	level.InitDefaults()
	return level
}
