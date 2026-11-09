package transport

import (
	"sync"
)

// ClientConnection represents an active connected client interface.
type ClientConnection interface {
	WriteMessage(message []byte) error
	Close()
}

// WebSocketHub manages active client connections and broadcasts messages.
type WebSocketHub struct {
	clients    map[ClientConnection]bool
	register   chan ClientConnection
	unregister chan ClientConnection
	broadcast  chan []byte
	mu         sync.Mutex
}

// NewWebSocketHub initializes a new WebSocketHub instance.
func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		clients:    make(map[ClientConnection]bool),
		register:   make(chan ClientConnection),
		unregister: make(chan ClientConnection),
		broadcast:  make(chan []byte),
	}
}

// Run starts the hub's main event loop for managing client registrations and broadcasting.
func (h *WebSocketHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				if err := client.WriteMessage(message); err != nil {
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// RegisterClient adds a client to the hub.
func (h *WebSocketHub) RegisterClient(client ClientConnection) {
	h.register <- client
}

// UnregisterClient removes and disconnects a client from the hub.
func (h *WebSocketHub) UnregisterClient(client ClientConnection) {
	h.unregister <- client
}

// BroadcastMessage sends a payload to all connected clients.
func (h *WebSocketHub) BroadcastMessage(message []byte) {
	h.broadcast <- message
}
