package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Turret struct {
	*Entity
	Direction Direction
	TimeSinceLastUpdate float64
}

type TurretBall struct {
	*Entity
	Speed     float64
	TimeSinceLastUpdate float64
}

func NewTurret(s *Stage, x, y int, direction Direction) *Turret {
	e := NewEntity(s, x, y, 1, 1, 'O', termbox.ColorMagenta, termbox.ColorBlack, nil, true)
	return &Turret{
		Entity:    e,
		Direction: direction,
	}
}

func NewTurretBall(s *Stage, t *Turret, x, y int, direction Direction) *TurretBall {
	e := NewEntity(s, x, y, 1, 1, '~', termbox.ColorMagenta, termbox.ColorBlack, nil, true)
	return &TurretBall{
		Entity:    e,
		Speed:     0.10,
	}
}

func (t *Turret) Update(s *Stage, event termbox.Event, delta time.Duration) {
	t.TimeSinceLastUpdate += delta.Seconds()

	if 2 > t.TimeSinceLastUpdate {
		return
	}

	t.OpenFire(s)
	t.TimeSinceLastUpdate = 0
}

func (t *TurretBall) Update(s *Stage, event termbox.Event, delta time.Duration) {
	t.TimeSinceLastUpdate += delta.Seconds()

	if 0.5 > t.TimeSinceLastUpdate {
		return
	}

	t.setPositionY(t.GetPositionY() + 1)
	t.TimeSinceLastUpdate = 0
}

func (t *Turret) OpenFire(s *Stage) {
	ball := NewTurretBall(s, t, t.GetPositionX(), t.GetPositionY() + 1, vertical)
	s.AddCanvasEntity(ball)
}
