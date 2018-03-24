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
	Stage  *Stage
	Logger *Logger
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
	stage.resize(termbox.Size())
	logger := NewLogger()
	game := &Game{stage, logger}
	stage.Game = game
	return game
}

type Renderer interface {
	Update(*Stage, termbox.Event, time.Duration)
	SetCells(*Stage)
}

// main game loop
// handles events, updates and renders stage and entities
func gameLoop(events chan termbox.Event, game *Game) *Game {
	termbox.Clear(fgColor, bgColor)
	stage := game.Stage
	stage.Render()
	lastUpdateTime := time.Now()

	for {
		termbox.Clear(fgColor, bgColor)
		update := time.Now()

		select {
		case key := <-events:
			switch {
			case key.Key == termbox.KeyCtrlC:
				// exit on ctrc + c
				return game
			default:
				stage.update(key, update.Sub(lastUpdateTime))
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

	_ = gameLoop(events, game)

	if len(game.Logger.logs) > 0 {
		game.Logger.DumpLogs()
		time.Sleep(2 * time.Second)
	}
	exit(events)
}
