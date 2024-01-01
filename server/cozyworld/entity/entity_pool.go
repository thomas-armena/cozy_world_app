package entity

import (
	"fmt"
)

type EntityPool struct {
	entities []Entity
}

func NewEntityPool() *EntityPool {
	return &EntityPool{
		entities: make([]Entity, 0),
	}
}

// TODO: Assign an ID when adding
func (e *EntityPool) AddEntity(ent Entity) {

	e.entities = append(e.entities, ent)
}

func Query[T any](e *EntityPool) []T {
	res := make([]T, 0)
	for _, maybeEnt := range e.entities {
		if ent, ok := maybeEnt.(T); ok {
			res = append(res, ent)
		}
	}
	return res
}

func QueryById[T any](e *EntityPool, id int32) (T, error) {
	// Empty value if not found. Is there a better approach here?
	var empt T

	for _, maybeEnt := range e.entities {
		if maybeEnt.EntityId() == id {
			ent, ok := maybeEnt.(T)
			if !ok {
				return empt, fmt.Errorf(
					"Found entity with id %v but has type mismatch: %v", id, maybeEnt)
			}
			return ent, nil
		}
	}

	return empt, fmt.Errorf("Could not find entity with id: %v", id)
}

// TODO: QueryByRegion (QuadTree), QueryNearest (QuadTree)
