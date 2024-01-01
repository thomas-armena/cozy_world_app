package entity

import (
	"github.com/go-gl/mathgl/mgl64"
	cpb "github.com/le-michael/cozyworld/protos"
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

func (m Motion) Start() mgl64.Vec2 {
	return m.start
}

func (m Motion) CurrentPosition(currUsec int64) mgl64.Vec2 {
	if m.start.ApproxEqual(m.dest) {
		return m.dest
	}

	distVec := m.dest.Sub(m.start)
	dirVec := distVec.Normalize()
	speedVec := dirVec.Mul(m.speed)

	deltaSec := float64((currUsec - m.startUsec)) / 1000000.0
	currPos := m.start.Add(speedVec.Mul(deltaSec))

	if distVec.LenSqr() < currPos.LenSqr() {
		return m.dest
	}
	return currPos
}

func (m Motion) MoveTo(currUsec int64, speed float64, dest mgl64.Vec2) Motion {
	currPos := m.CurrentPosition(currUsec)
	return Motion{
		start:     currPos,
		startUsec: currUsec,
		speed:     speed,
		dest:      dest,
	}
}

func (m Motion) ToProto() *cpb.Motion {
	return &cpb.Motion{
		Start:     &cpb.Vec2{X: m.start.X(), Y: m.start.Y()},
		StartUsec: m.startUsec,
		Speed:     m.speed,
		Dest:      &cpb.Vec2{X: m.dest.X(), Y: m.dest.Y()},
	}
}

func (m *Motion) FromProto(pb *cpb.Motion) {
	m.start = mgl64.Vec2{pb.Start.X, pb.Start.Y}
	m.startUsec = pb.StartUsec
	m.speed = pb.Speed
	m.dest = mgl64.Vec2{pb.Dest.X, pb.Dest.Y}
}
