package fantasia

//import "github.com/nsf/termbox-go"

const Level1TileMapString = `
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
                        |        |        |     |     |     |
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
|        |     |        |     |              |               
+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
`

func NewLevel1(g *Game) *Level {
	// create user
	user := NewUser(g.Stage, 1, 1)
	var entities []Renderer
	entities = append(entities, user)

	// create title
	title := NewWord(g.Stage, 0, 3, "Level 1")
	explanation := NewWord(g.Stage, 0, 5, "J: down, H: left, K: up, L: right")
	g.Stage.AddScreenEntity(title, explanation)

	return &Level{
		Game:          g,
		Entities:      entities,
		TileMapString: Level1TileMapString,
	}
}
