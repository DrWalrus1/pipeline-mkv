package makemkv

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func (h *MakeMkvRouteHandler) BackupHandler(w http.ResponseWriter, r *http.Request) {
	decrypt, err := strconv.ParseBool(r.URL.Query().Get("decrypt"))
	if err != nil {
		r.Response.StatusCode = 400
		return
	}
	source := r.URL.Query().Get("source")
	destination := r.URL.Query().Get("destination")

	conn, err := h.SocketHandler.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	messageChan := h.SocketHandler.ReadClientMessages(conn)

	updates, cancel, err := h.CommandHandler.TriggerDiskBackup(decrypt, source, destination)
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
