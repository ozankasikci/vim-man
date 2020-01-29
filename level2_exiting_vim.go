package vimman

import "github.com/nsf/termbox-go"

const LevelExitingVimTileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                            VIM                           |
|~                  Hardest editor to exit                  |
|~                            :)                            |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
|~                                                          |
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
			CollisionCallback:  func() {},
		},
	}

	level := &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: LevelExitingVimTileMapString,
		TileData:      tileData,
		InputBlocked:  true,
		VimMode:       normalMode,
		Init: func() {
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 2 - EXITING VIM", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "You can't be a great Vim user without knowing how to exit.", explanationOptions)

			hintOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg, CenterHorizontally: true}
			hint := NewWord(g.Stage, levelHintCoordX, levelHintCoordY, "Type colon ':', then 'q', press enter", hintOptions)

			g.Stage.AddScreenEntity(title, explanation, hint)
		},
		ColonLineCallbacks: make(map[string]func(*Game)),
	}

	exitTerms := []string{"q", "quit", "exit"}
	for _, term := range exitTerms {
		if _, ok := level.ColonLineCallbacks[term]; !ok {
			level.ColonLineCallbacks[term] = func(g *Game) {
				levelInstance := NewLevelTextEditing(g)
				g.Stage.SetLevel(levelInstance)
			}
		}
	}

	level.InitDefaults()
	return level
}
