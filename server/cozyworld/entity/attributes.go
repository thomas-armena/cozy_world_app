package entity

import (
	"github.com/go-gl/mathgl/mgl64"

	cpb "github.com/le-michael/cozyworld/protos"
)

type Entity interface {
	EntityId() int32
	AssignEntityId(id int32)
	ToProto() *cpb.Entity
}

type Position interface {
	CurrentPosition()
}

type Movable interface {
	MoveTo(currUsec int64, dest mgl64.Vec2)
}
