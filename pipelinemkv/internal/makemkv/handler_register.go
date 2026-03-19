package makemkv

import (
	"fmt"
	"net/http"
	"strings"

	streamtracker "github.com/DrWalrus1/pipelinemkv/internal/streamTracker"
	"github.com/DrWalrus1/pipelinemkv/internal/websocket"
)

func SetupMkvCommandApiPaths(mux *http.ServeMux, handler gomakemkv.MakeMkvCommandHandler, tracker streamtracker.StreamTracker, socketHandler websocket.WebSocketHandler) {
	mux.HandleFunc("/api/info", GetDiskInfoHandler(handler, socketHandler))
	mux.HandleFunc("/api/mkv", GetSaveDiskInfoHandler(handler, &tracker, socketHandler))
	mux.HandleFunc("/api/watch/mkv", GetWatchMkvHandler(&tracker, socketHandler))
	mux.HandleFunc("/api/backup", GetBackupHandler(handler, &tracker, socketHandler))
	mux.HandleFunc("POST /api/register", GetRegisterHandler(handler))
}

func validateSource(source string) error {
	if source == "" {
		return fmt.Errorf("source cannot be empty")
	}

	if strings.HasPrefix(source, "disc:") || strings.HasPrefix(source, "iso:") || strings.HasPrefix(source, "file:") || strings.HasPrefix(source, "dev:") {
		return nil
	}
	return fmt.Errorf("invalid source")
}
