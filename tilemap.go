package vimman

import (
	"github.com/nsf/termbox-go"
	"strings"
)

func NewTileMapCell(ch rune, fn func(), lineNumber int) TileMapCellData {
	return TileMapCellData{
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 ch,
		CollidesPhysically: false,
		CollisionCallback:  fn,
		LineNumber:         lineNumber,
	}
}

type TileMapCellData struct {
	Ch                 rune
	BgColor            termbox.Attribute
	FgColor            termbox.Attribute
	CollidesPhysically bool
	CollisionCallback  func()
	InitCallback       func(*Entity)
	LineNumber         int
}

type TileMapCellDataMap map[rune]TileMapCellData

var CommonTileMapCellData = TileMapCellDataMap{
	'0': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 ' ',
		CollidesPhysically: false,
		CollisionCallback:  func() {},
	},
	'↓': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '↓',
		CollidesPhysically: false,
		CollisionCallback:  func() {},
	},
	'+': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '+',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'-': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '-',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'|': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '|',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'█': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '█',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'◼': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '◼',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'▅': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '▅',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'▀': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '▀',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'☵': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '☵',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	'☲': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 '☲',
		CollidesPhysically: true,
		CollisionCallback:  func() {},
	},
	' ': {
		BgColor:            termbox.ColorBlack,
		FgColor:            termbox.ColorWhite,
		Ch:                 ' ',
		CollidesPhysically: false,
		CollisionCallback:  func() {},
	},
}

func ParseLine(l string) []rune {
	var lineChars []rune

	//chars := strings.Split(l, " ")
	//line := strings.Join(chars, "")

	for _, char := range l {
		lineChars = append(lineChars, char)
	}

	return lineChars
}

func ParseTileMapString(tileMap string) [][]rune {
	var parsed [][]rune

	lines := strings.Split(tileMap, "\n")
	lines = lines[1 : len(lines)-1]

	for _, line := range lines {
		l := ParseLine(line)
		parsed = append(parsed, l)
	}

	return parsed
}
