package fantasia

import "github.com/nsf/termbox-go"

type TileMapCell struct {
	*termbox.Cell
}

type TileMapCellData map[rune]struct{
	ch rune
	bgColor termbox.Attribute
	fgColor termbox.Attribute
}

