package fantasia

import "github.com/nsf/termbox-go"

//import "github.com/nsf/termbox-go"

const level3TileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
Level 2                 |        |        |     |     |     |
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
|        |     |        |     |              |              ↓ 
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
`

func NewLevel3(g *Game) *Level {
	// create user
	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)

	// create title
	titleOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
	title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 3 - Vim Modes", titleOptions)

	explanationOptions := WordOptions{InitCallback: nil, Fg: levelTitleFg, Bg: levelTitleBg}
	explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "i: Insert Mode, v: Visual Mode", explanationOptions)
	g.Stage.AddScreenEntity(title, explanation)

	tileData := TileMapCellDataMap{
		'↓': TileMapCellData{
			ch:                 '↓',
			fgColor:            termbox.ColorGreen,
			bgColor:            termbox.ColorBlack,
			collidesPhysically: false,
			collisionCallback: func() {

			},
		},
	}

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: level3TileMapString,
		TileData:      tileData,
	}
}
