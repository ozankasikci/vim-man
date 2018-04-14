package fantasia

const tileMapString = `
1 + 1 1 1 1 1
1 0 1 0 0 0 1
1 0 1 0 1 0 1
1 0 1 0 1 0 1
1 0 1 0 1 0 1
1 0 0 0 1 0 1
1 1 1 1 1 - 1
`
func NewLevel1(s *Stage) *Level {
	user := NewUser()
	var entities []Renderer
	entities = append(entities, user)

	return &Level{
		Entities: entities,
		TileMapString: tileMapString,
	}
}
