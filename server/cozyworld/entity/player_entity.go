package entity

import (
	"github.com/go-gl/mathgl/mgl64"
	cpb "github.com/le-michael/cozyworld/protos"
)

type PlayerEntity struct {
	Entity
	Movable

	entityId int32
	motion   Motion
}

func NewPlayerEntity(currUsec int64) *PlayerEntity {
	// TODO: Don't hard code spawn point
	return &PlayerEntity{
		motion: NewIdleMotion(mgl64.Vec2{1.2, 1.2}, currUsec),
	}
}

func (p *PlayerEntity) EntityId() int32 {
	return p.entityId
}

func (p *PlayerEntity) AssignEntityId(id int32) {
	p.entityId = id
}

func (p *PlayerEntity) MoveTo(currUsec int64, dest mgl64.Vec2) {
	// TODO: speed variable?
	p.motion = p.motion.MoveTo(currUsec, 4, dest)
}

func (p *PlayerEntity) ToProto() *cpb.Entity {
	return &cpb.Entity{
		EntityId: p.entityId,
		Entity: &cpb.Entity_PlayerEntity_{
			PlayerEntity: &cpb.Entity_PlayerEntity{
				Motion: p.motion.ToProto(),
			},
		},
	}
}
