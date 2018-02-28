package fantasia

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

const (
	bgColor = termbox.ColorBlack
	fgColor = termbox.ColorYellow
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
	Stage  *Stage
	Logger *Logger
}

type GameOptions struct {
	fps          float64
	initialLevel int
	bgCell       *termbox.Cell
}

func NewGame(opts GameOptions) *Game {
	bgCell := &termbox.Cell{'~', termbox.ColorGreen, termbox.ColorBlack}
	stage := NewStage(opts.initialLevel, opts.fps, bgCell)
	stage.Init()
	stage.resize(termbox.Size())
	logger := NewLogger()
	game := &Game{stage, logger}
	return game
}

type Renderer interface {
	Update(*Stage, termbox.Event)
	SetCells(*Stage)
}

func gameLoop(events chan termbox.Event, game *Game) *Game {
	termbox.Clear(fgColor, bgColor)
	stage := game.Stage
	stage.render()

	for {
		termbox.Clear(fgColor, bgColor)
		update := time.Now()

		select {
		case key := <-events:
			switch {
			case key.Key == termbox.KeyCtrlC:
				return game
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
	termbox.Clear(termbox.ColorDefault, bgColor)

	events := make(chan termbox.Event)
	go eventLoop(events)

	game := NewGame(GameOptions{
		fps:          80,
		initialLevel: 1,
	})

	_ = gameLoop(events, game)

	if len(game.Logger.logs) > 0 {
		game.Logger.DumpLogs()
		time.Sleep(2 * time.Second)
	}
	exit(events)
}
