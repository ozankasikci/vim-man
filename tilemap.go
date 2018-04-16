package fantasia

import "github.com/nsf/termbox-go"

type TileMapCell struct {
	*termbox.Cell
}

type TileMapCellData map[rune]struct {
	ch      rune
	bgColor termbox.Attribute
	fgColor termbox.Attribute
}

var CommonTileMapCellData = TileMapCellData{
	'1': {
		bgColor: termbox.ColorRed,
		fgColor: termbox.ColorRed,
		ch:      '░',
	},
	'0': {
		bgColor: termbox.ColorBlack,
		fgColor: termbox.ColorWhite,
		ch:      ' ',
	},
	'+': {
		bgColor: termbox.ColorYellow,
		fgColor: termbox.ColorYellow,
		ch:      ' ',
	},
	'-': {
		bgColor: termbox.ColorBlue,
		fgColor: termbox.ColorBlue,
		ch:      ' ',
	},
}
