package entity

import (
	"github.com/go-gl/mathgl/mgl64"
)

// A Motion can be used to define a Moveable entity with a Position.
type Motion struct {
	start     mgl64.Vec2
	startUsec int64
	// Speed of the entity in m/sec
	speed float64
	dest  mgl64.Vec2
}

func NewIdleMotion(pos mgl64.Vec2, currUsec int64) Motion {
	return NewMotion(pos, currUsec, 0, pos)
}

func NewMotion(start mgl64.Vec2, startUsec int64, speed float64, dest mgl64.Vec2) Motion {
	return Motion{
		start:     start,
		startUsec: startUsec,
		speed:     speed,
		dest:      dest,
	}
}

func (p Motion) Start() mgl64.Vec2 {
	return p.start
}

func (p Motion) CurrentPosition(currUsec int64) mgl64.Vec2 {
	if p.start.ApproxEqual(p.dest) {
		return p.dest
	}

	distVec := p.dest.Sub(p.start)
	dirVec := distVec.Normalize()
	speedVec := dirVec.Mul(p.speed)

	deltaSec := float64((currUsec - p.startUsec)) / 1000000.0
	currPos := p.start.Add(speedVec.Mul(deltaSec))

	if distVec.LenSqr() < currPos.LenSqr() {
		return p.dest
	}
	return currPos
}

func (p Motion) MoveTo(currUsec int64, speed float64, dest mgl64.Vec2) Motion {
	currPos := p.CurrentPosition(currUsec)
	return Motion{
		start:     currPos,
		startUsec: currUsec,
		speed:     speed,
		dest:      dest,
	}
}
