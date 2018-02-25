package fantasia

import (
	"github.com/nsf/termbox-go"
	"github.com/pkg/errors"
	"time"
)

const (
	bgColor = termbox.ColorYellow
	fgColor = termbox.ColorBlack
)

type point struct {
	x int
	y int
}

type word struct {
	text     string
	location point
	cursor   int
}

type Game struct {
	Stage *Stage
}

type Renderer interface {
	Update(*Stage, termbox.Event)
	Render(*Stage)
}

func gameLoop(events chan termbox.Event, stage *Stage) *Stage {
	termbox.Clear(fgColor, bgColor)
	stage.render()

	for {
		termbox.Clear(fgColor, bgColor)
		update := time.Now()

		select {
		case key := <-events:
			switch {
			case key.Key == termbox.KeyEsc:
				return stage
			default:
				stage.update(key)
			}
		}

		stage.render()
		time.Sleep(time.Duration((update.Sub(time.Now()).Seconds()*1000.0)+1000.0/stage.Fps) * time.Millisecond)
	}
}

func eventLoop(e chan termbox.Event) {
	for {
		e <- termbox.PollEvent()
	}
}

func exit(events chan termbox.Event) {
	close(events)
	termbox.Close()
}

func Init() {
	if err := termbox.Init(); err != nil {
		panic(errors.Wrap(err, "failed to init termbox"))
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)

	events := make(chan termbox.Event)
	go eventLoop(events)

	stage := NewStage(1, 100)
	stage.Init()
	stage.resize(termbox.Size())

	_ = gameLoop(events, stage)
	exit(events)
}
