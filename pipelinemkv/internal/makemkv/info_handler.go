package makemkv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/gorilla/websocket"
)

type diskInfoHandler interface {
	TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error)
}

type webSocketHandler interface {
	GetUpgrader() *websocket.Upgrader
	ReadClientMessages(conn *websocket.Conn) <-chan string
	SendClientUpdates(conn *websocket.Conn, updates <-chan []byte, clientMessages <-chan string, clientMessageHandler func(string) bool)
}

func GetDiskInfoHandler(h diskInfoHandler, websocketHandler webSocketHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.URL.Query().Get("source")
		conn, err := websocketHandler.GetUpgrader().Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		done := websocketHandler.ReadClientMessages(conn)

		reader, cancel, err := h.TriggerDiskInfo(source)
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
		//TODO: add function to log updates if needed
		_, discInfoChan := gomakemkv.ParseMakeMkvInfoCommandLogs(reader)
		updatesInBytes := make(chan []byte)
		//TODO: FIX ME
		discInfoInBytes, err := json.Marshal(discInfoChan)
		updatesInBytes <- discInfoInBytes

		clientMessageHandler := func(message string) bool {
			// do nothing because we automatically close when channel closes
			return false
		}
		websocketHandler.SendClientUpdates(conn, updatesInBytes, done, clientMessageHandler)
	}
}
