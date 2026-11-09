package transport

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Conn implements the ClientConnection interface wrapping a WebSocket connection.
type Conn struct {
	ws  *websocket.Conn
	hub *WebSocketHub
}

// WriteMessage sends a binary message to the WebSocket client.
func (c *Conn) WriteMessage(message []byte) error {
	return c.ws.WriteMessage(websocket.BinaryMessage, message)
}

// Close terminates the underlying WebSocket connection.
func (c *Conn) Close() {
	c.ws.Close()
}

// HandleWebSocket upgrades HTTP requests to WebSockets and handles the client read/write lifecycle.
func HandleWebSocket(hub *WebSocketHub, w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade WebSocket connection: %v\n", err)
		return
	}

	client := &Conn{
		ws:  wsConn,
		hub: hub,
	}

	hub.RegisterClient(client)
	router := NewRouter()

	// Read loop running in its own goroutine for each connected client
	go func() {
		defer func() {
			hub.UnregisterClient(client)
		}()

		for {
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected WebSocket close error: %v\n", err)
				}
				break
			}

			// Route the incoming raw payload
			router.RouteMessage(client, message)
		}
	}()
}
