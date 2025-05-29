package routehandlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"pipelinemkv/makemkv"
	makemkvCommands "pipelinemkv/makemkv/commands"
	osCommands "pipelinemkv/os/commands"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development).  **SECURITY WARNING**:  In production, restrict this!
	},
}

type RouteHandler struct {
	StreamTracker *makemkv.StreamTracker
}

func readClientMessages(conn *websocket.Conn) <-chan string {
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

func (handler *RouteHandler) sendClientUpdates(conn *websocket.Conn, updates <-chan []byte, clientMessages <-chan string, clientMessageHandler func(string) bool) {
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

func (handler *RouteHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	done := readClientMessages(conn)

	reader, cancel, err := makemkvCommands.TriggerDiskInfo(source)
	defer cancel()
	if err != nil {
		log.Printf("Could not trigger get disk info: %v", err)
		err = conn.WriteMessage(websocket.TextMessage, fmt.Appendf(nil, "Could not trigger get disk info: %v", err))
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
		return
	}
	updates := makemkvCommands.WatchInfoLogs(reader)

	clientMessageHandler := func(message string) bool {
		// do nothing because we automatically close when channel closes
		return false
	}
	handler.sendClientUpdates(conn, updates, done, clientMessageHandler)
}

func (handler *RouteHandler) MkvHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	title := r.URL.Query().Get("title")
	destination := r.URL.Query().Get("destination")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	messageChan := readClientMessages(conn)

	reader, cancel, err := makemkvCommands.TriggerSaveMkv(source, title, destination)
	handler.StreamTracker.AddStream(source, &reader, cancel)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not trigger makemkv save: %v", err)
		log.Println(errorMessage)
		err = conn.WriteMessage(websocket.TextMessage, []byte(errorMessage))
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
		return
	}
	updates := makemkvCommands.WatchSaveMkvLogs(reader)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			handler.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	handler.sendClientUpdates(conn, updates, messageChan, clientMessageHandler)
}

func (handler *RouteHandler) WatchMkv(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")

	reader, ok := handler.StreamTracker.GetStream(source)
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("No live stream found"))
		return
	}
	cancel := handler.StreamTracker.GetStreamCancelFunc(source)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	messageChan := readClientMessages(conn)

	updates := makemkvCommands.WatchSaveMkvLogs(*reader)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			handler.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	handler.sendClientUpdates(conn, updates, messageChan, clientMessageHandler)
}

func (handler *RouteHandler) BackupHandler(w http.ResponseWriter, r *http.Request) {
	decrypt, err := strconv.ParseBool(r.URL.Query().Get("decrypt"))
	if err != nil {
		r.Response.StatusCode = 400
		return
	}
	source := r.URL.Query().Get("source")
	destination := r.URL.Query().Get("destination")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	messageChan := readClientMessages(conn)

	reader, cancel, err := makemkvCommands.TriggerDiskBackup(decrypt, source, destination)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not trigger disk backup: %v", err)
		log.Println(errorMessage)
		err = conn.WriteMessage(websocket.TextMessage, []byte(errorMessage))
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
		return
	}

	updates := makemkvCommands.WatchBackupLogs(reader)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			handler.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	handler.sendClientUpdates(conn, updates, messageChan, clientMessageHandler)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	responseStatus := makemkvCommands.RegisterMkvKey(key)
	w.WriteHeader(responseStatus)
}

func EjectHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	log.Printf("Ejecting device: %s", device)

	responseStatus := osCommands.EjectDevice(device)
	if responseStatus != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not eject device: " + responseStatus.Error()))
		return
	}
	w.WriteHeader(200)
}

func InsertDiscHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	log.Printf("Inserting device: %s", device)

	responseStatus := osCommands.InsertDevice(device)
	if responseStatus != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not insert device: " + responseStatus.Error()))
		return
	}
	w.WriteHeader(200)
}
