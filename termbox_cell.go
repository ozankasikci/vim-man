package fantasia

import "github.com/nsf/termbox-go"

type TermBoxCell struct {
	*termbox.Cell
	collidesPhysically bool
	cellData           TileMapCellData
}
