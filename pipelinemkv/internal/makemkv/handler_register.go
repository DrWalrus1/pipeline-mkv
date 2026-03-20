package makemkv

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/DrWalrus1/gomakemkv"
	streamtracker "github.com/DrWalrus1/pipelinemkv/pipelinemkv/internal/streamTracker"
	"github.com/DrWalrus1/pipelinemkv/pipelinemkv/internal/websocket"
)

func SetupMkvCommandApiPaths(mux *http.ServeMux, handler MakeMkvRouteHandler, tracker streamtracker.StreamTracker, socketHandler websocket.WebSocketHandler) {
	mux.HandleFunc("/api/info", handler.DiskInfoHandler)
	mux.HandleFunc("/api/mkv", handler.SaveDiskInfoHandler)
	mux.HandleFunc("/api/watch/mkv", handler.WatchMkvHandler)
	mux.HandleFunc("/api/backup", handler.BackupHandler)
	mux.HandleFunc("POST /api/register", handler.RegisterHandler)
}

type MakeMkvRouteHandler struct {
	CommandHandler *gomakemkv.MakeMkvCommandHandler
	SocketHandler  *websocket.WebSocketHandler
	StreamTracker  *streamtracker.StreamTracker
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
