package fantasia

import (
	"github.com/nsf/termbox-go"
	"strings"
	"time"
)

type Level struct {
	Game     *Game
	TileMapString  string
	TileMap  [][]TileMapCell
	TileData TileMapCellData
	Entities []Renderer
	BgCell   *termbox.Cell
	Width    int
	Height   int
}

func parseLine(l string) []rune {
	var lineChars []rune

	chars := strings.Split(l, " ")
	line := strings.Join(chars, "")

	for _, char := range line {
		lineChars = append(lineChars, char)
	}

	return lineChars
}

func parseTileMapString(tileMap string) [][]rune {
	var parsed[][]rune

	lines := strings.Split(tileMap, "\n")
	lines = lines[1:len(lines) - 1]

	for _, line := range lines {
		l := parseLine(line)
		parsed = append(parsed, l)
	}

	return parsed
}

func (l *Level) Update(s *Stage, t time.Duration) {

}

func (l *Level) SetCells(s *Stage) {

}

func (l *Level) LoadTileMapCells(parsedRunes [][]rune) [][]TileMapCell {
	var cells [][]TileMapCell

	for _, line := range parsedRunes {
		rowCells := make([]TileMapCell, len(line))

		for j, char := range line {
			if data, ok := l.TileData[char]; !ok {
				panic("Couldn't retrieve tile data")
			} else {
				cell := TileMapCell{ &termbox.Cell{ data.ch, data.fgColor, data.bgColor } }
				rowCells[j] = cell
			}
		}
		for k := 0; k < len(rowCells); k++ {
			//lg.LogValue(rowCells[k].Ch, rowCells[k].Bg)
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

