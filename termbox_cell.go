package fantasia

import "github.com/nsf/termbox-go"

type TermBoxCell struct {
	*termbox.Cell
	collidesPhysically bool
	cellData           TileMapCellData
}

func EmptyTileMapCell() *TermBoxCell {
	data := CommonTileMapCellData[' ']
	cell := &TermBoxCell{
		&termbox.Cell{data.Ch, data.FgColor, data.BgColor},
		data.CollidesPhysically,
		data,
	}

	return cell
}
