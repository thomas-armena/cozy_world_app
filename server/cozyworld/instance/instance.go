package instance

import (
	"log"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/le-michael/cozyworld/client"
	"github.com/le-michael/cozyworld/clock"
	"github.com/le-michael/cozyworld/entity"
	"github.com/le-michael/cozyworld/quadtree"
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

type InstanceAction func(i *Instance)

type Instance struct {
	clock clock.Clock

	nextEntityId int32

	closed      bool
	actionQueue chan InstanceAction

	clients []client.Client

	bounds   quadtree.Rect
	entities map[int32]entity.Entity
}

func NewInstance(clock clock.Clock, top, left, width, height float64) *Instance {
	bounds := quadtree.NewClosedRect(top, left, width, height)

	return &Instance{
		clock: clock,

		nextEntityId: 0,

		actionQueue: make(chan InstanceAction),

		bounds:   bounds,
		entities: make(map[int32]entity.Entity),
	}
}

func (i *Instance) NextEntityId() int32 {
	id := i.nextEntityId
	i.nextEntityId++
	return id
}

func (i *Instance) AddEntity(e entity.Entity) int32 {
	e.AssignEntityId(i.NextEntityId())
	i.entities[e.EntityId()] = e
	return e.EntityId()
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

func (i *Instance) AddClient(c client.Client) {
	i.clients = append(i.clients, c)
	p := entity.NewPlayerEntity(i.clock.CurrentTime())
	i.AddEntity(p)
	c.AssignEntityId(p.EntityId())

	i.BroadcastUpdateEntity(c.EntityId())
}

func (i *Instance) Close() {
	i.closed = true
	// Pass in a noop to end the run loop incase it is empty.
	i.actionQueue <- func(i *Instance) {}
	close(i.actionQueue)
}

func (i *Instance) HandleMoveToCommand(entityId int32, mc *cpb.InstanceStreamRequest_MoveToCommand) {
	// Handle missing
	e := i.entities[entityId]

	// Handle failed casting
	m := e.(entity.Movable)

	m.MoveTo(i.clock.CurrentTime(), mgl64.Vec2{mc.Dest.X, mc.Dest.Y})

	i.BroadcastUpdateEntity(entityId)
}

func (i *Instance) BroadcastUpdateEntity(id int32) {
	e := i.entities[id]

	res := &cpb.InstanceStreamResponse{
		Command: &cpb.InstanceStreamResponse_UpdateEntityCommand_{
			UpdateEntityCommand: &cpb.InstanceStreamResponse_UpdateEntityCommand{
				Entity: e.ToProto(),
			},
		},
	}
	log.Printf("Broadcasting update entity command: %v\n", res.String())

	i.Broadcast(res)
}

func (i *Instance) Broadcast(response *cpb.InstanceStreamResponse) {
	data, err := proto.Marshal(response)
	if err != nil {
		log.Printf("Failed to parse response: %v\n", err)
	}
	for _, c := range i.clients {
		c.Write(data)
	}
}
