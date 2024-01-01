package client

import (
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

type FakeClient struct {
	Client

	entityId      int32
	writeCallback func([]byte)
}

func NewFakeClient(writeCallback func([]byte)) *FakeClient {
	return &FakeClient{
		writeCallback: writeCallback,
	}
}

func (f *FakeClient) Write(data []byte) {
	f.writeCallback(data)
}

func (f *FakeClient) AssignEntityId(entityId int32) {
	// TODO make this generic
	f.entityId = entityId

	res := &cpb.InstanceStreamResponse{Command: &cpb.InstanceStreamResponse_ConnectionCommand_{
		ConnectionCommand: &cpb.InstanceStreamResponse_ConnectionCommand{
			EntityId: entityId,
		},
	}}
	data, _ := proto.Marshal(res) // Handle error?

	f.Write(data)
}

func (f *FakeClient) EntityId() int32 {
	return f.entityId
}
