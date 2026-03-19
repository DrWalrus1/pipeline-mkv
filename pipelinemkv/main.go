package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/DrWalrus1/pipelinemkv/internal/config"
	makeMkvHandlers "github.com/DrWalrus1/pipelinemkv/internal/makemkv"
	"github.com/DrWalrus1/pipelinemkv/internal/metadata/tmdb"
	"github.com/DrWalrus1/pipelinemkv/internal/optical"
	st "github.com/DrWalrus1/pipelinemkv/internal/streamTracker"
	"github.com/DrWalrus1/pipelinemkv/internal/websocket"
)

//go:embed static/*
var vueFiles embed.FS

func main() {

	var configPath string
	flag.StringVar(&configPath, "config", "", "filepath for config.json file")
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}
	commandHandler := gomakemkv.MakeMkvCommandHandler{
		ExecutablePath: conf.ExecutablePath,
	}

	meta_service := tmdb.New(conf.MetadataServiceToken)
	meta_service.SearchMovie(context.Background(), "Forrest Gump", "", "")
	meta_service.GetMovieDetails(context.Background(), "550", []string{})

	websocketHandler := websocket.NewHandler()
	streamTracker := st.NewStreamTracker()
	mux := http.NewServeMux()
	makeMkvHandlers.SetupMkvCommandApiPaths(mux, commandHandler, streamTracker, websocketHandler)
	config.SetupConfigApiPaths(mux, conf)
	optical.SetupOpticalApiPaths(mux)

	disFS, _ := fs.Sub(vueFiles, "static")
	mux.HandleFunc("/", serveStaticFrontend(disFS))

	fmt.Printf("WebSocket server started on %s\n", conf.Port)
	err = http.ListenAndServe(conf.Port, mux)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func serveStaticFrontend(disFS fs.FS) func(w http.ResponseWriter, r *http.Request) {
	staticServer := http.FileServer(http.FS(disFS))

	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			http.Error(w, "API endpoint not found", http.StatusNotFound)
			return
		}

		// FIX 3: Check if the file exists in the embedded FS
		// We trim the leading slash to match the FS root
		f, err := disFS.Open(strings.TrimPrefix(r.URL.Path, "/"))
		if err != nil {
			// File doesn't exist (like a Vue route /dashboard)
			// Serve index.html to allow SPA routing to take over
			index, err := disFS.Open("index.html")
			if err != nil {
				http.Error(w, "Index not found", http.StatusNotFound)
				return
			}
			defer index.Close()

			// Just copy the index file out
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			io.Copy(w, index)
			return
		}
		f.Close()

		// File exists, let the static server handle it
		staticServer.ServeHTTP(w, r)

	}
}
