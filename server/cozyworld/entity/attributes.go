package entity

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Entity interface {
	EntityId() int64
	AssignEntityId(id int64)
}

type Position interface {
	CurrentPosition()
}

type Movable interface {
	MoveTo(dest mgl64.Vec2)
}
