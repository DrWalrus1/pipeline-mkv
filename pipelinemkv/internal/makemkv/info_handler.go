package makemkv

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type diskInfoHandler interface {
	TriggerDiskInfo(source string) (io.Reader, context.CancelFunc, error)
}

func (h *MakeMkvRouteHandler) DiskInfoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	conn, err := h.SocketHandler.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	done := h.SocketHandler.ReadClientMessages(conn)

	discInfoChan, cancel, err := h.CommandHandler.TriggerDiskInfo(source, context.Background())
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
	updatesInBytes := make(chan []byte)
	//TODO: FIX ME
	discInfoInBytes, err := json.Marshal(discInfoChan)
	updatesInBytes <- discInfoInBytes

	clientMessageHandler := func(message string) bool {
		// do nothing because we automatically close when channel closes
		return false
	}
	h.SocketHandler.SendClientUpdates(conn, updatesInBytes, done, clientMessageHandler)
}
