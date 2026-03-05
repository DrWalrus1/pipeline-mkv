package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"pipelinemkv/cmd/makemkv"
	st "pipelinemkv/cmd/streamTracker"
	"pipelinemkv/config"
	"pipelinemkv/routehandlers"
	metadataservice "pipelinemkv/services/metadata_service"
)

var commandHandler makemkv.IMakeMkvCommandHandler

func main() {
	if commandHandler == nil {
		commandHandler = &makemkv.MakeMkvCommandHandler{}
	}

	var configPath string
	flag.StringVar(&configPath, "config", "", "filepath for config.json file")
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	meta_service := metadataservice.New(conf.MetadataServiceToken)
	meta_service.SearchMovie(context.Background(), "Forrest Gump", "", "")
	meta_service.GetMovieDetails(context.Background(), "550", []string{})

	streamTracker := st.NewStreamTracker()
	advancedHandler := routehandlers.RouteHandler{
		StreamTracker:  &streamTracker,
		MakeMkvHandler: commandHandler,
	}
	mux := http.NewServeMux()
	SetupPaths(mux, advancedHandler)

	fs := http.FileServer(http.Dir("./static/"))
	handler := ServeWithoutHTMLExtension{fs: fs, staticFolder: "./static/"}
	mux.HandleFunc("/", handler.ServerHTTP)

	fmt.Printf("WebSocket server started on %s\n", conf.Port)
	err = http.ListenAndServe(conf.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func SetupPaths(mux *http.ServeMux, advancedHandler routehandlers.RouteHandler) {
	http.HandleFunc("/api/info", advancedHandler.InfoHandler)
	http.HandleFunc("/api/mkv", advancedHandler.MkvHandler)
	http.HandleFunc("/api/watch/mkv", advancedHandler.WatchMkv)
	http.HandleFunc("/api/backup", advancedHandler.BackupHandler)
	http.HandleFunc("POST /api/register", advancedHandler.RegistrationHandler)
	http.HandleFunc("POST /api/eject", routehandlers.EjectHandler)
	http.HandleFunc("POST /api/insert", routehandlers.InsertDiscHandler)
}
