package instance

import (
	"github.com/le-michael/cozyworld/client"
	"github.com/le-michael/cozyworld/entity"
	"github.com/le-michael/cozyworld/quadtree"
)

type InstanceAction func(i *Instance)

type Instance struct {
	entityId int32

	closed      bool
	actionQueue chan InstanceAction

	clients []client.Client

	bounds   quadtree.Rect
	entities *quadtree.QuadTreeNode[entity.Entity]
}

func NewInstance(top, left, width, height float64) *Instance {
	bounds := quadtree.NewClosedRect(top, left, width, height)

	return &Instance{
		entityId: 0,

		actionQueue: make(chan InstanceAction),

		bounds:   bounds,
		entities: quadtree.NewQuadTreeNode[entity.Entity](bounds, nil),
	}
}

func (i *Instance) NextEntityId() int32 {
	id := i.entityId
	i.entityId++
	return id
}

// ScheduleAction adds an action to the queue. Returns true if it is successful.
func (i *Instance) ScheduleAction(action InstanceAction) bool {
	if i.closed {
		return false
	}

	i.actionQueue <- action
	return true
}

func (i *Instance) Run() {
	for action := range i.actionQueue {
		if i.closed {
			break
		}
		action(i)
	}
}

func (i *Instance) AddClient(c client.Client) int32 {
	i.clients = append(i.clients, c)
	return i.NextEntityId()
}

func (i *Instance) Close() {
	i.closed = true
	// Pass in a noop to end the run loop incase it is empty.
	i.actionQueue <- func(i *Instance) {}
	close(i.actionQueue)
}
