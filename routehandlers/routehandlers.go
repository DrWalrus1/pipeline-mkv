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

func (handler *RouteHandler) InfoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// TODO: add error handling
	reader, cancel, err := makemkvCommands.TriggerDiskInfo(source)
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
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				cancel()
				return
			}
			fmt.Printf("Message Type: %v", messageType)
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					cancel()
					return
				}
				return
			}
		}
	}()
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			cancel()
			return // Exit if we can't write (client likely disconnected)
		}
	}
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

	reader, cancel, err := makemkvCommands.TriggerSaveMkv(source, title, destination)
	handler.StreamTracker.AddStream(source, &reader)
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
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				handler.StreamTracker.RemoveStream(source)
				cancel()
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
			}
		}
	}()
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
}

func (handler *RouteHandler) WatchMkv(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")

	reader, ok := handler.StreamTracker.GetStream(source)
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("No live stream found"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	updates := makemkvCommands.WatchSaveMkvLogs(*reader)
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				handler.StreamTracker.RemoveStream(source)
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
			}
		}
	}()
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
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

	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				cancel()
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
			}
		}
	}()
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
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
