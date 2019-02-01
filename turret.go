package fantasia

import (
	"github.com/nsf/termbox-go"
	"time"
)

type Turret struct {
	*Entity
	Direction Direction
	TurretBallSpeed float64
	TimeSinceLastUpdate float64
}

type TurretBall struct {
	*Entity
	Speed     float64
	TimeSinceLastUpdate float64
}

func NewTurret(s *Stage, x, y int, direction Direction, turretBallSpeed float64) *Turret {
	e := NewEntity(s, x, y, 1, 1, 'O', termbox.ColorMagenta, termbox.ColorBlack, nil, true)
	return &Turret{
		Entity:    e,
		Direction: direction,
		TurretBallSpeed: turretBallSpeed,
	}
}

func NewTurretBall(s *Stage, t *Turret, x, y int, direction Direction, speed float64) *TurretBall {
	e := NewEntity(s, x, y, 1, 1, '~', termbox.ColorMagenta, termbox.ColorBlack, nil, true)

	if speed == 0 {
        GetLogger().Log("noo")
		speed = 0.10
	}
	GetLogger().LogValue(speed)

	return &TurretBall{
		Entity: e,
		Speed:  speed,
	}
}

func (t *Turret) Update(s *Stage, event termbox.Event, delta time.Duration) {
	t.TimeSinceLastUpdate += delta.Seconds()

	if 2 > t.TimeSinceLastUpdate {
		return
	}

	t.OpenFire(s, t.TurretBallSpeed)
	t.TimeSinceLastUpdate = 0
}

func (t *TurretBall) Update(s *Stage, event termbox.Event, delta time.Duration) {
	t.TimeSinceLastUpdate += delta.Seconds()

	if 0.5 > t.TimeSinceLastUpdate {
		return
	}


	t.setPositionY(t.GetPositionY() + int(t.Speed))
	t.TimeSinceLastUpdate = 0
}

func (t *Turret) OpenFire(s *Stage, speed float64) {
	ball := NewTurretBall(s, t, t.GetPositionX(), t.GetPositionY() + 1, vertical, speed)
	s.AddCanvasEntity(ball)
}
