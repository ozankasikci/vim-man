package fantasia

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

const (
	bgColor = termbox.ColorBlack
	fgColor = termbox.ColorWhite
	KeyB    = 98
	KeyE    = 101
	KeyH    = 104
	KeyJ    = 106
	KeyK    = 107
	KeyL    = 108
)

type point struct {
	x int
	y int
}

type Game struct {
	Stage *Stage
	screenSizeX int
	screenSizeY int
}

type GameOptions struct {
	fps          float64
	initialLevel int
	bgCell       *termbox.Cell
}

func NewGame(opts GameOptions) *Game {
	bgCell := &termbox.Cell{'░', fgColor, bgColor}
	stage := NewStage(opts.initialLevel, opts.fps, bgCell)
	stage.Init()
	game := &Game{ stage, 0, 0 }
	stage.Game = game
	return game
}

type Renderer interface {
	Update(*Stage, termbox.Event, time.Duration)
	SetCells(*Stage)
}

// main game loop
// handles events, updates and renders stage and entities
func gameLoop(events chan termbox.Event, game *Game) {
	termbox.Clear(fgColor, bgColor)
	stage := game.Stage
	game.setScreenSize(termbox.Size())
	stage.Render()
	lastUpdateTime := time.Now()

	for {
		termbox.Clear(fgColor, bgColor)
		update := time.Now()

		select {
		case event := <-events:
			switch {
			case event.Key == termbox.KeyCtrlC:
				// exit on ctrc + c
				return
			case event.Type == termbox.EventResize:
				game.setScreenSize(termbox.Size())
			default:
				stage.update(event, update.Sub(lastUpdateTime))
			}
		default:
			stage.update(termbox.Event{}, update.Sub(lastUpdateTime))
		}
		lastUpdateTime = time.Now()

		stage.Render()
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
	termbox.Clear(termbox.ColorDefault, bgColor)

	events := make(chan termbox.Event)
	go eventLoop(events)

	game := NewGame(GameOptions{
		fps:          50,
		initialLevel: 1,
	})

	gameLoop(events, game)

	if len(lg.logs) > 0 {
		lg.DumpLogs()
		time.Sleep(2 * time.Second)
	}

	exit(events)
}

func (g *Game) setScreenSize(x, y int) {
	lg.LogValue(x, y)
	if x > 0 {
		g.screenSizeX = x
	}

	if y > 0 {
		g.screenSizeY = y
	}
}

func (g *Game) getScreenSize() (int, int) {
	return g.screenSizeX, g.screenSizeY
}
