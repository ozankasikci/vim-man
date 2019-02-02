package fantasia

import "github.com/nsf/termbox-go"

//import "github.com/nsf/termbox-go"

const Level2TileMapString = `
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

func NewLevel2(g *Game) *Level {
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
				levelInstance := NewLevel3(g)
				g.Stage.SetLevel(levelInstance)
			},
		},
	}

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: Level1TileMapString,
		TileData:      tileData,
		Init: func() {
			// load info
			title := NewWord(g.Stage, levelTitleCoordX, levelTitleCoordY, "Level 2 - Vim Modes")
			explanation := NewWord(g.Stage, levelExplanationCoordX, levelExplanationCoordY, "i: Insert Mode, v: Visual Mode")
			g.Stage.AddScreenEntity(title, explanation)
		},
	}
}
