package fantasia

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type TermBoxCell struct {
	*termbox.Cell
	collidesPhysically bool
	cellData           TileMapCellData
}

func NewTileMapCell(ch rune, fn func()) TileMapCellData {
	return TileMapCellData{
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 ch,
		collidesPhysically: false,
		collisionCallback:  fn,
	}
}

type TileMapCellData struct {
	ch                 rune
	bgColor            termbox.Attribute
	fgColor            termbox.Attribute
	collidesPhysically bool
	collisionCallback  func()
	initCallback       func(*Entity)
}

type TileMapCellDataMap map[rune]TileMapCellData

var CommonTileMapCellData = TileMapCellDataMap{
	'1': {
		bgColor:            termbox.ColorWhite,
		fgColor:            termbox.ColorWhite,
		ch:                 ' ',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'0': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 ' ',
		collidesPhysically: false,
		collisionCallback:  func() {},
	},
	'↓': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '↓',
		collidesPhysically: false,
		collisionCallback:  func() {},
	},
	'+': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '+',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'-': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '-',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'|': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '|',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'█': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '█',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'◼': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '◼',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'▅': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '▅',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'▀': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '▀',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'☵': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '☵',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	'☲': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 '☲',
		collidesPhysically: true,
		collisionCallback:  func() {},
	},
	' ': {
		bgColor:            termbox.ColorBlack,
		fgColor:            termbox.ColorWhite,
		ch:                 ' ',
		collidesPhysically: false,
		collisionCallback:  func() {},
	},
}

func parseLine(l string) []rune {
	var lineChars []rune

	//chars := strings.Split(l, " ")
	//line := strings.Join(chars, "")

	for _, char := range l {
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
