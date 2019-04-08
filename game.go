package fantasia

import (
	"time"

	"github.com/nsf/termbox-go"
	"github.com/pkg/errors"
)

const (
	bgColor = termbox.ColorBlack
	fgColor = termbox.ColorWhite
)

type Point struct {
	x int
	y int
}

type Game struct {
	Stage       *Stage
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
	game := &Game{nil, 0, 0}
	stage := NewStage(game, opts.initialLevel, opts.fps, bgCell)
	game.Stage = stage
	return game
}

type Renderer interface {
	Update(*Stage, termbox.Event, time.Duration)
	Destroy()
	SetCells(*Stage)
	GetCells() []*TermBoxCell
	GetPosition() (int, int)
	GetPositionX() int
	GetPositionY() int
	GetScreenOffset() (int, int)
	GetDrawPriority() int
	GetTags() []Tag
	ShouldCenterHorizontally() bool
}

// main game loop
// handles events, updates and renders stage and entities
func gameLoop(events chan termbox.Event, game *Game) {
	termbox.Clear(fgColor, bgColor)
	game.setScreenSize(termbox.Size())
	stage := game.Stage
	stage.Init()
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
	termbox.Clear(termbox.ColorDefault, bgColor)

	events := make(chan termbox.Event)
	go eventLoop(events)

	game := NewGame(GameOptions{
		fps:          50,
		initialLevel: 1,
	})

	// main game loop, this is blocking
	gameLoop(events, game)

	// dump logs after the gameLoop stops
	if len(lg.logs) > 0 {
		lg.DumpLogs()
		time.Sleep(2 * time.Second)
	}

	exit(events)
}

func (g *Game) setScreenSize(x, y int) {
	if x > 0 {
		g.screenSizeX = x
	}

	if y > 0 {
		g.screenSizeY = y
	}
}

func (g *Game) getScreenSize() (int, int) {
	return g.getScreenSizeX(), g.getScreenSizeY()
}

func (g *Game) getScreenSizeX() int {
	return g.screenSizeX
}

func (g *Game) getScreenSizeY() int {
	return g.screenSizeY
}
