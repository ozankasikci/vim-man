package fantasia

import "github.com/nsf/termbox-go"

type TileMapCell struct {
	*termbox.Cell
	def TileMapDef
}

type TileMapDef map[string]struct{
	ch string
	bgColor termbox.Attribute
}

