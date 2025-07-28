package routehandlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pipelinemkv/cmd/makemkv"
	streamtracker "pipelinemkv/cmd/streamTracker"
	osCommands "pipelinemkv/os/commands"
	"strconv"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/DrWalrus1/gomakemkv/events"
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
	MakeMkvHandler makemkv.IMakeMkvCommandHandler
	StreamTracker  *streamtracker.StreamTracker
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

func (h *RouteHandler) sendClientUpdates(conn *websocket.Conn, updates <-chan []byte, clientMessages <-chan string, clientMessageHandler func(string) bool) {
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

func (h *RouteHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	done := readClientMessages(conn)

	reader, cancel, err := h.MakeMkvHandler.TriggerDiskInfo(source)
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
	updates, discInfoChan := gomakemkv.ParseMakeMkvInfoCommandLogs(reader)
	updatesInBytes := make(chan []byte)
	go func() {
		isUpdatesComplete := false
		isDiscInfoReceived := false
		for {
			if isUpdatesComplete && isDiscInfoReceived {
				return
			}
			select {
			case update, ok := <-updates:
				if !ok {
					isUpdatesComplete = true
					break
				}
				updateInBytes, _ := json.Marshal(update)
				updatesInBytes <- updateInBytes
			case discInfo, ok := <-discInfoChan:
				if !ok {
					isDiscInfoReceived = true
					break
				}
				updateInBytes, _ := json.Marshal(discInfo)
				updatesInBytes <- updateInBytes
			}
		}
	}()

	clientMessageHandler := func(message string) bool {
		// do nothing because we automatically close when channel closes
		return false
	}
	h.sendClientUpdates(conn, updatesInBytes, done, clientMessageHandler)
}

func (h *RouteHandler) MkvHandler(w http.ResponseWriter, r *http.Request) {
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

	reader, cancel, err := h.MakeMkvHandler.TriggerSaveMkv(source, title, destination)
	h.StreamTracker.AddStream(source, &reader, cancel)
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
	updates := gomakemkv.ParseMakeMkvLogs(reader)
	updatesInBytes := stringifyMakeMkvOutput(updates)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			h.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	h.sendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)
}

func (h *RouteHandler) WatchMkv(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")

	reader, ok := h.StreamTracker.GetStream(source)
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("No live stream found"))
		return
	}
	cancel := h.StreamTracker.GetStreamCancelFunc(source)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	messageChan := readClientMessages(conn)

	updates := gomakemkv.ParseMakeMkvLogs(*reader)
	updatesInBytes := stringifyMakeMkvOutput(updates)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			h.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	h.sendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)
}

func (h *RouteHandler) BackupHandler(w http.ResponseWriter, r *http.Request) {
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

	reader, cancel, err := h.MakeMkvHandler.TriggerDiskBackup(decrypt, source, destination)
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

	updates := gomakemkv.ParseMakeMkvLogs(reader)
	updatesInBytes := stringifyMakeMkvOutput(updates)

	clientMessageHandler := func(message string) bool {
		if message == "cancel" {
			h.StreamTracker.RemoveStream(source)
			cancel()
			return true
		}
		return false
	}

	h.sendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)
}

func (h *RouteHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	err := h.MakeMkvHandler.RegisterMakeMkv(key)
	if err != nil {
		if err == gomakemkv.ErrBadKey {
			w.WriteHeader(http.StatusBadRequest)
		} else if err == gomakemkv.ErrUnexpectedRegistrationError {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
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

func stringifyMakeMkvOutput(updates <-chan events.MakeMkvOutput) chan []byte {
	updatesInBytes := make(chan []byte)
	go func() {
		for {
			select {
			case update, ok := <-updates:
				if !ok {
					return
				}
				marshalled, err := json.Marshal(update)
				if err != nil {
					log.Println("Failed to convert update to json")
					continue
				}
				updatesInBytes <- marshalled
			}
		}
	}()
	return updatesInBytes
}
