package makemkv

import (
	"log"
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
)

func GetWatchMkvHandler(tracker StreamTracker, sh webSocketHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.URL.Query().Get("source")

		reader, ok := tracker.GetStream(source)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte("No live stream found"))
			return
		}
		cancel := tracker.GetStreamCancelFunc(source)

		conn, err := sh.GetUpgrader().Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		messageChan := sh.ReadClientMessages(conn)

		updates := gomakemkv.ParseMakeMkvLogs(*reader)
		updatesInBytes := stringifyMakeMkvOutput(updates)

		clientMessageHandler := func(message string) bool {
			if message == "cancel" {
				tracker.RemoveStream(source)
				cancel()
				return true
			}
			return false
		}

		sh.SendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)

	}
}
