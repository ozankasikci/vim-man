package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Level struct {
	Game          *Game
	TileMapString string
	TileMap       [][]*TileMapCell
	TileData      TileMapCellDataMap
	Entities      []Renderer
	BgCell        *termbox.Cell
	Width         int
	Height        int
}

func (l *Level) Update(s *Stage, t time.Duration) {

}

func (l *Level) SetCells(s *Stage) {

}

func (l *Level) GetSize() (int, int) {
	return len(l.TileMap[0]), len(l.TileMap)
}

func (l *Level) GetScreenOffset() (int, int) {
	offsetX, offsetY := 0, 0
	screenWidth, screenHeight := l.Game.getScreenSize()
	levelWidth, levelHeight := l.GetSize()

	if screenWidth > levelWidth {
		offsetX = (screenWidth - levelWidth) / 2
	}

	if screenHeight > levelHeight {
		offsetY = (screenHeight - levelHeight) / 2
	}

	return offsetX, offsetY
}

func (l *Level) LoadTileMapCells(parsedRunes [][]rune) [][]*TileMapCell {
	var cells [][]*TileMapCell

	for _, line := range parsedRunes {
		rowCells := make([]*TileMapCell, len(line))

		for j, char := range line {
			if data, ok := l.TileData[char]; !ok {
				panic("Couldn't retrieve tile data")
			} else {
				cell := &TileMapCell{
					&termbox.Cell{data.ch, data.fgColor, data.bgColor},
					data.collides,
				}
				rowCells[j] = cell
			}
		}

		cells = append(cells, rowCells)
	}

	l.TileMap = cells
	return l.TileMap
}

func (l *Level) LoadTileMap() {
	parsed := parseTileMapString(l.TileMapString)
	l.LoadTileMapCells(parsed)
}

// row, length
func (l *Level) GetTileMapDimensions() (int, int) {
	parsed := parseTileMapString(l.TileMapString)
	rowLength := len(parsed[0])
	columnLength := len(parsed)
	return rowLength, columnLength
}
