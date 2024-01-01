package client

import (
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

type Client interface {
	Write([]byte)
	AssignEntityId(int32)

	EntityId() int32
}

type WebsocketClient struct {
	Client

	entityId int32
	conn     net.Conn
}

func NewWebsocketClient(conn net.Conn) *WebsocketClient {
	return &WebsocketClient{conn: conn}
}

func (w *WebsocketClient) Write(data []byte) {
	// Handle error
	wsutil.WriteServerMessage(w.conn, ws.OpBinary, data)
}

func (w *WebsocketClient) AssignEntityId(entityId int32) {
	w.entityId = entityId 

	res := &cpb.InstanceStreamResponse{Command: &cpb.InstanceStreamResponse_ConnectionCommand_{
		ConnectionCommand: &cpb.InstanceStreamResponse_ConnectionCommand{
			EntityId: entityId,
		},
	}}
	data, _ := proto.Marshal(res) // Handle error?

	w.Write(data)
}

func (w *WebsocketClient) EntityId() int32 {
	return w.entityId
}
