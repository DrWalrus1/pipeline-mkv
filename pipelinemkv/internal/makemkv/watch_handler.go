package makemkv

import (
	"log"
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
)

func (h *MakeMkvRouteHandler) WatchMkvHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")

	reader, ok := h.StreamTracker.GetStream(source)
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("No live stream found"))
		return
	}
	cancel := h.StreamTracker.GetStreamCancelFunc(source)

	conn, err := h.SocketHandler.GetUpgrader().Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	messageChan := h.SocketHandler.ReadClientMessages(conn)

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

	h.SocketHandler.SendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)

}
