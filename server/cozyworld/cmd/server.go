package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/le-michael/cozyworld/client"
	"github.com/le-michael/cozyworld/clock"
	"github.com/le-michael/cozyworld/instance"
	"google.golang.org/protobuf/proto"

	cpb "github.com/le-michael/cozyworld/protos"
)

const (
	port = 8080
)

var world = instance.NewInstance(clock.NewSystemClock(), 0, 0, 5, 5)

func handleWsConnection(conn net.Conn) {
	defer conn.Close()

	client := client.NewWebsocketClient(conn)
	world.AddClient(client)

	state := ws.StateServerSide
	r := wsutil.NewReader(conn, state)

	for {
		if _, err := r.NextFrame(); err != nil {
			log.Printf("Unable to read next frame: %v", err)
			return
		}

		data, err := io.ReadAll(r)
		if err != nil {
			log.Printf("Unable to read next frame: %v", err)
			return
		}

		var req cpb.InstanceStreamRequest
		if err := proto.Unmarshal(data, &req); err != nil {
			log.Printf("Failed to Unmarshal request proto: %v", err)
			return
		}

		// TODO: Handle more commands
		mc := req.GetMoveToCommand()
		world.HandleMoveToCommand(client.EntityId(), mc)

		log.Printf("Request: %v\n", req.String())
	}
}

func main() {
	log.Printf("Starting a server on port: %v", port)

	// TODO: Handle multiple instances
	go world.Run()
	defer world.Close()

	err := http.ListenAndServe(
		fmt.Sprintf(":%v", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			conn, _, _, err := ws.UpgradeHTTP(r, w)
			if err != nil {
				log.Printf("Unabled to upgrade http connection to ws: %v", err)
				return
			}
			log.Printf("Upgraded a new http client!")

			go handleWsConnection(conn)
		}))
	if err != nil {
		log.Fatalf("Unable to start server: %v\n", err)
	}
}
