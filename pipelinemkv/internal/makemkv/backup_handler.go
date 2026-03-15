package makemkv

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/gorilla/websocket"
)

type backupHandler interface {
	TriggerDiskBackup(decrypt bool, source string, destination string) (io.Reader, context.CancelFunc, error)
}

func GetBackupHandler(h backupHandler, tracker StreamTracker, wsh webSocketHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decrypt, err := strconv.ParseBool(r.URL.Query().Get("decrypt"))
		if err != nil {
			r.Response.StatusCode = 400
			return
		}
		source := r.URL.Query().Get("source")
		destination := r.URL.Query().Get("destination")

		conn, err := wsh.GetUpgrader().Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		messageChan := wsh.ReadClientMessages(conn)

		reader, cancel, err := h.TriggerDiskBackup(decrypt, source, destination)
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
				tracker.RemoveStream(source)
				cancel()
				return true
			}
			return false
		}

		wsh.SendClientUpdates(conn, updatesInBytes, messageChan, clientMessageHandler)
	}
}
