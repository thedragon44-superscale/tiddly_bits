package transport

import (
	"log"

	"emv_system/server/handlers"
)

// Router inspects incoming client payloads and dispatches them to the appropriate subsystem handlers.
type Router struct{}

// NewRouter initializes and returns a new Router instance.
func NewRouter() *Router {
	return &Router{}
}

// RouteMessage determines the message type based on the leading opcode byte and routes it accordingly.
func (rt *Router) RouteMessage(client ClientConnection, payload []byte) {
	if len(payload) == 0 {
		log.Println("Received empty payload from client")
		return
	}

	// The first byte acts as the command opcode/routing selector
	opcode := payload[0]
	data := payload[1:]

	switch opcode {
	case 0x01:
		log.Println("Routing frame to APDU handler")
		handlers.HandleAPDUCommand(client, data)

	case 0x02:
		log.Println("Routing frame to Query handler")
		handlers.HandleQuery(client, data)

	default:
		log.Printf("Unrecognized message opcode: 0x%x\n", opcode)
	}
}
