package fantasia

import (
	"github.com/nsf/termbox-go"
	"reflect"
	"time"
)

type VimMode int

const (
	normalMode VimMode = iota
	insertMode
)

const (
	levelTitleCoordX       int = 0
	levelTitleCoordY       int = 1
	levelTitleFg           termbox.Attribute = termbox.ColorGreen
	levelTitleBg           termbox.Attribute = termbox.ColorBlack
	levelExplanationCoordX int = 0
	levelExplanationCoordY int = 2
	typedCharacterFg       termbox.Attribute = termbox.ColorWhite
	typedCharacterBg       termbox.Attribute = termbox.ColorBlack
)

type Level struct {
	Game          *Game
	VimMode       VimMode
	TileMapString string
	TileMap       [][]*TermBoxCell
	TileData      TileMapCellDataMap
	Entities      []Renderer
	BgCell        *termbox.Cell
	Width         int
	Height        int
	Init          func()
}

func (l *Level) Update(s *Stage, t time.Duration) {

}

func (l *Level) SetCells(s *Stage) {

}

func (l *Level) GetSize() (int, int) {
	index, length := 0, 0

	for i, line := range l.TileMap {
		if len(line) > length {
			index, length = i, len(line)
		}
	}

	return len(l.TileMap[index]), len(l.TileMap)
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

func (l *Level) LoadTileMapCells(parsedRunes [][]rune) [][]*TermBoxCell {
	var cells [][]*TermBoxCell

	for _, line := range parsedRunes {
		rowCells := make([]*TermBoxCell, len(line))
		var data TileMapCellData

		for j, char := range line {
			if _, ok := l.TileData[char]; !ok {
				if _, ok := CommonTileMapCellData[char]; !ok {
					data = NewTileMapCell(char, func() {})
				} else {
					data = CommonTileMapCellData[char]
				}
			} else {
				data = l.TileData[char]
			}

			if reflect.DeepEqual(data, TileMapCellData{}) {
				data = CommonTileMapCellData[char]
			}

			cell := &TermBoxCell{
				&termbox.Cell{data.ch, data.fgColor, data.bgColor},
				data.collidesPhysically,
				data,
			}
			rowCells[j] = cell
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
