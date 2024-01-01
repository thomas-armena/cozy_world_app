package instance_test

import (
	"fmt"
	"testing"

	"github.com/le-michael/cozyworld/client"
	"github.com/le-michael/cozyworld/clock"
	"github.com/le-michael/cozyworld/instance"
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

func TestAddClient(t *testing.T) {
	responses := []*cpb.InstanceStreamResponse{}
	c := client.NewFakeClient(func(b []byte) {
		var msg cpb.InstanceStreamResponse
		proto.Unmarshal(b, &msg)
		responses = append(responses, &msg)
	})

	world := instance.NewInstance(clock.NewFakeClock(0), 0, 0, 5, 5)
	world.AddClient(c)

	expectedResponses := []*cpb.InstanceStreamResponse{
		// Connection command
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_ConnectionCommand_{
				ConnectionCommand: &cpb.InstanceStreamResponse_ConnectionCommand{
					EntityId: 0,
				},
			},
		},
		// Update entity broadcast
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_UpdateEntityCommand_{
				UpdateEntityCommand: &cpb.InstanceStreamResponse_UpdateEntityCommand{
					Entity: &cpb.Entity{
						EntityId: 0,
						Entity: &cpb.Entity_PlayerEntity_{
							PlayerEntity: &cpb.Entity_PlayerEntity{
								Motion: &cpb.Motion{
									Start:     &cpb.Vec2{X: 1.2, Y: 1.2},
									StartUsec: 0,
									Speed:     0,
									Dest:      &cpb.Vec2{X: 1.2, Y: 1.2},
								},
							},
						},
					},
				},
			},
		},
	}
	if err := expectResponsesMatch(responses, expectedResponses); err != nil {
		t.Fatal(err)
	}
}

func TestMultipleClientConnections(t *testing.T) {
	responses1 := []*cpb.InstanceStreamResponse{}
	responses2 := []*cpb.InstanceStreamResponse{}

	c1 := client.NewFakeClient(func(b []byte) {
		var msg cpb.InstanceStreamResponse
		proto.Unmarshal(b, &msg)
		responses1 = append(responses1, &msg)
	})
	c2 := client.NewFakeClient(func(b []byte) {
		var msg cpb.InstanceStreamResponse
		proto.Unmarshal(b, &msg)
		responses2 = append(responses2, &msg)
	})

	world := instance.NewInstance(clock.NewFakeClock(0), 0, 0, 5, 5)
	world.AddClient(c1)
	world.AddClient(c2)

	if c1.EntityId() != 0 {
		t.Fatalf("c1.EntityId() = %v; want = %v", c1.EntityId(), 0)
	}
	if c2.EntityId() != 1 {
		t.Fatalf("c2.EntityId() = %v; want = %v", c2.EntityId(), 1)
	}

	expectedResponses1 := []*cpb.InstanceStreamResponse{
		// Connection command
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_ConnectionCommand_{
				ConnectionCommand: &cpb.InstanceStreamResponse_ConnectionCommand{
					EntityId: 0,
				},
			},
		},
		// Update entity broadcast
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_UpdateEntityCommand_{
				UpdateEntityCommand: &cpb.InstanceStreamResponse_UpdateEntityCommand{
					Entity: &cpb.Entity{
						EntityId: 0,
						Entity: &cpb.Entity_PlayerEntity_{
							PlayerEntity: &cpb.Entity_PlayerEntity{
								Motion: &cpb.Motion{
									Start:     &cpb.Vec2{X: 1.2, Y: 1.2},
									StartUsec: 0,
									Speed:     0,
									Dest:      &cpb.Vec2{X: 1.2, Y: 1.2},
								},
							},
						},
					},
				},
			},
		},
		// Update entity broadcast
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_UpdateEntityCommand_{
				UpdateEntityCommand: &cpb.InstanceStreamResponse_UpdateEntityCommand{
					Entity: &cpb.Entity{
						EntityId: 1,
						Entity: &cpb.Entity_PlayerEntity_{
							PlayerEntity: &cpb.Entity_PlayerEntity{
								Motion: &cpb.Motion{
									Start:     &cpb.Vec2{X: 1.2, Y: 1.2},
									StartUsec: 0,
									Speed:     0,
									Dest:      &cpb.Vec2{X: 1.2, Y: 1.2},
								},
							},
						},
					},
				},
			},
		},
	}
	if err := expectResponsesMatch(responses1, expectedResponses1); err != nil {
		t.Fatal(err)
	}

	expectedResponses2 := []*cpb.InstanceStreamResponse{
		// Connection command with client 1's entity sent back as well
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_ConnectionCommand_{
				ConnectionCommand: &cpb.InstanceStreamResponse_ConnectionCommand{
					EntityId: 0,
					Entities: []*cpb.Entity{
						&cpb.Entity{
							EntityId: 0,
							Entity: &cpb.Entity_PlayerEntity_{
								PlayerEntity: &cpb.Entity_PlayerEntity{
									Motion: &cpb.Motion{
										Start:     &cpb.Vec2{X: 1.2, Y: 1.2},
										StartUsec: 0,
										Speed:     0,
										Dest:      &cpb.Vec2{X: 1.2, Y: 1.2},
									},
								},
							},
						},
					},
				},
			},
		},
		// Update entity broadcast
		&cpb.InstanceStreamResponse{
			Command: &cpb.InstanceStreamResponse_UpdateEntityCommand_{
				UpdateEntityCommand: &cpb.InstanceStreamResponse_UpdateEntityCommand{
					Entity: &cpb.Entity{
						EntityId: 1,
						Entity: &cpb.Entity_PlayerEntity_{
							PlayerEntity: &cpb.Entity_PlayerEntity{
								Motion: &cpb.Motion{
									Start:     &cpb.Vec2{X: 1.2, Y: 1.2},
									StartUsec: 0,
									Speed:     0,
									Dest:      &cpb.Vec2{X: 1.2, Y: 1.2},
								},
							},
						},
					},
				},
			},
		},
	}
	if err := expectResponsesMatch(responses2, expectedResponses2); err != nil {
		t.Fatal(err)
	}
}

func expectResponsesMatch(actual, expected []*cpb.InstanceStreamResponse) error {
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"Responses len do not match. got = %v; want = %v", len(actual), len(expected))
	}

	for i := range expected {
		if !proto.Equal(actual[i], expected[i]) {
			fmt.Errorf(
				"Unexpected response.\n\t got = %v;\n\t want = %v", actual[i].String(), expected[i].String())
		}
	}

	return nil
}
