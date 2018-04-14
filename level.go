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

func loadTileMapCells(parsedRunes [][]rune)  {
	lg.LogValue(parsedRunes)
}

func (l *Level) LoadTileMap() {
	parsed := parseTileMapString(l.TileMapString)
	loadTileMapCells(parsed)
}

// row, length
func (l *Level) GetTileMapDimensions() (int, int) {
	parsed := parseTileMapString(l.TileMapString)
	rowLength := len(parsed[0])
	columnLength := len(parsed)
	return rowLength, columnLength
}

