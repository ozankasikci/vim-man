package fantasia

//import "github.com/nsf/termbox-go"

const Level1TileMapString = `
1 + 1 1 1 1 1
1 0 1 0 0 0 1
1 0 1 0 1 0 1
1 0 1 0 1 0 1
1 0 1 0 1 0 1
1 0 0 0 1 0 1
1 1 1 1 1 - 1
`

var Level1TileData = TileMapCellData{
	'1': CommonTileMapCellData['1'],
	'0': CommonTileMapCellData['0'],
	'+': CommonTileMapCellData['+'],
	'-': CommonTileMapCellData['-'],
}

func NewLevel1() *Level {
	user := NewUser()
	var entities []Renderer
	entities = append(entities, user)

	return &Level{
		Entities:      entities,
		TileMapString: Level1TileMapString,
		TileData:      Level1TileData,
	}
}
