package websocket

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	upgrader websocket.Upgrader
}

func NewHandler() WebSocketHandler {
	return WebSocketHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins (for development).  **SECURITY WARNING**:  In production, restrict this!
			},
		},
	}
}

func (h WebSocketHandler) GetUpgrader() *websocket.Upgrader {
	return &h.upgrader
}

func (h WebSocketHandler) ReadClientMessages(conn *websocket.Conn) <-chan string {
	done := make(chan string)

	go func() {
		defer close(done) // Signal completion when this goroutine exits
		for {
			// ReadMessage blocks until a message is received or an error occurs.
			// We don't care about the message content here, just the connection status.
			_, message, err := conn.ReadMessage()
			if err != nil {
				// Check if the error indicates a normal client disconnect.
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					log.Println("Client disconnected (detected by read pump).")
				} else {
					log.Printf("WebSocket read error: %v", err)
				}
				return // Exit the goroutine on any read error or disconnect
			}
			done <- string(message)
			// If you had a need to process incoming messages from the client (e.g., pings, control messages),
			// you would do so here. For this handler, we are only sending data to the client.
		}
	}()
	return done
}

func (h WebSocketHandler) SendClientUpdates(conn *websocket.Conn, updates <-chan []byte, clientMessages <-chan string, clientMessageHandler func(string) bool) {
	for {
		select {
		case update, ok := <-updates:
			if !ok {
				return
			}
			err := conn.WriteMessage(websocket.TextMessage, update)
			if err != nil {
				log.Println("write error:", err)
				return // Exit if we can't write (client likely disconnected)
			}
		case message, ok := <-clientMessages:
			if !ok {
				return
			}
			if clientMessageHandler(message) {
				return
			}
		}
	}
}
