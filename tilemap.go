package fantasia

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type TileMapCell struct {
	*termbox.Cell
	collides bool
}

func (t *TileMapCell) GetCellData() TileMapCellData {
	return CommonTileMapCellData[t.Ch]
}

type TileMapCellData struct {
	ch       rune
	bgColor  termbox.Attribute
	fgColor  termbox.Attribute
	collides bool
}

type TileMapCellDataMap map[rune]TileMapCellData

var CommonTileMapCellData = TileMapCellDataMap{
	'1': {
		bgColor:  termbox.ColorWhite,
		fgColor:  termbox.ColorRed,
		ch:       '░',
		collides: true,
	},
	'0': {
		bgColor:  termbox.ColorBlack,
		fgColor:  termbox.ColorWhite,
		ch:       ' ',
		collides: false,
	},
	'↓': {
		bgColor:  termbox.ColorBlack,
		fgColor:  termbox.ColorWhite,
		ch:       '↓',
		collides: false,
	},
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
	var parsed [][]rune

	lines := strings.Split(tileMap, "\n")
	lines = lines[1 : len(lines)-1]

	for _, line := range lines {
		l := parseLine(line)
		parsed = append(parsed, l)
	}

	return parsed
}
