package fantasia

import "github.com/nsf/termbox-go"

//import "github.com/nsf/termbox-go"

const level1TileMapString = `
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

func NewLevel1(g *Game) *Level {
	// create user
	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)

	tileData := TileMapCellDataMap{
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
		TileMapString: level1TileMapString,
		TileData:      tileData,
		Init: func() {
			// load info
			titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 1 - Basic Movement", titleOptions)

			explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "J: down, H: left, K: up, L: right", explanationOptions)

			g.Stage.AddScreenEntity(title, explanation)
		},
	}
}
