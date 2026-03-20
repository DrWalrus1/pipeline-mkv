package makemkv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/DrWalrus1/gomakemkv/events"
	"github.com/gorilla/websocket"
)

type discSaveHandler interface {
	TriggerSaveMkv(source string, title string, destination string) (io.Reader, context.CancelFunc, error)
}

type StreamTracker interface {
	GetStream(key string) (*io.Reader, bool)
	GetStreamCancelFunc(key string) context.CancelFunc
	AddStream(key string, reader *io.Reader, cancelFunc context.CancelFunc) error
	RemoveStream(key string)
}

func (h *MakeMkvRouteHandler) SaveDiskInfoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	title := r.URL.Query().Get("title")
	destination := r.URL.Query().Get("destination")

	conn, err := h.SocketHandler.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	messageChan := h.SocketHandler.ReadClientMessages(conn)

	reader, cancel, err := h.CommandHandler.TriggerSaveMkv(source, title, destination)
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

	h.SocketHandler.SendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)
}

func stringifyMakeMkvOutput(updates <-chan events.MakeMkvOutput) chan []byte {
	updatesInBytes := make(chan []byte)
	go func() {
		for update := range updates {
			marshalled, err := json.Marshal(update)
			if err != nil {
				log.Println("Failed to convert update to json")
				continue
			}
			updatesInBytes <- marshalled
		}
	}()
	return updatesInBytes
}
