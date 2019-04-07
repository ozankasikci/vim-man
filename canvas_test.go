package fantasia

import (
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanvasCheckCollision(t *testing.T) {
	x, y := 10, 10
	c := NewCanvas(x, y)
	c[1][1] = &TermBoxCell{&termbox.Cell{}, false}
	c[0][0] = &TermBoxCell{&termbox.Cell{}, true}

	tt := []struct {
		x      int
		y      int
		expect bool
	}{
		{-1, 0, true},
		{0, -1, true},
		{1, 1, false},
		{0, 0, true},
		{x, 0, true},
		{0, y, true},
	}

	for _, value := range tt {
		res := c.CheckCollision(value.x, value.y)
		assert.Equal(t, value.expect, res)
	}

	c[2][2] = nil
	assert.True(t, c.CheckCollision(2, 2))
}
