package handlers

import (
	"log"
)

// ClientConnection defines the minimal required methods to communicate with a client.
type ClientConnection interface {
	WriteMessage(message []byte) error
}

// HandleAPDUCommand processes incoming APDU command frames from the client.
func HandleAPDUCommand(client ClientConnection, data []byte) {
	log.Printf("Received APDU command payload of length: %d bytes\n", len(data))

	// TODO: Bridge to engine/reader to execute APDU command against card reader session,
	// then serialize and return the response payload.

	// Placeholder response (e.g., Status Word 0x9000 indicating success)
	response := []byte{0x90, 0x00}
	if err := client.WriteMessage(response); err != nil {
		log.Printf("Failed to write APDU response to client: %v\n", err)
	}
}
