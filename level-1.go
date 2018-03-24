package fantasia

const tileMapString = `
1211111
1212221
1212121
1222121
1111121
`

type Level1 struct {
	*Level
	Stage *Stage
}

func NewLevel1() *Level{
	user := NewUser()
	var entities []Renderer
	entities = append(entities, user)

	return &Level{
		Entities: entities,
		Width: 10,
		Height: 5,
	}
}

