package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Level struct {
	Entities  []Renderer
	BgCell    *termbox.Cell
	Width     int
	Height    int
}

func (l *Level) Update(s *Stage, t time.Duration)  {

}

func (l *Level) SetCells(s *Stage)  {

}

func (l *Level ) LoadTileMap(tileMap string)  {
	type Data struct {
		Votes *Votes `json:"votes"`
		Count string `json:"count,omitempty"`
	}

	type Votes struct {
		OptionA string `json:"option_A"`
	}
}
