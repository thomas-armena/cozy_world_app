package client

import (
	"log"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

type Client interface {
	Write(data []byte)
	AssignEntityId(id int32)
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

func (w *WebsocketClient) AssignEntityId(id int32) {
	log.Printf("Assigning entity id: %v\n", id)
	w.entityId = id

	command := &cpb.InstanceStreamResponse_ConnectionCommand{
		EntityId: id,
	}
	data, _ := proto.Marshal(command) // Handle error?

	w.Write(data)
}
