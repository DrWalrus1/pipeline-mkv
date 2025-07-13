package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"pipelinemkv/cmd/makemkv"
	st "pipelinemkv/cmd/streamTracker"
	"pipelinemkv/routehandlers"
	"time"
)

func main() {
	var port string
	flag.StringVar(&port, "port", ":9090", "Port to host the server on")
	flag.Parse()
	// runInitialDiscLoadOnStartup()
	streamTracker := st.NewStreamTracker()
	advancedHandler := routehandlers.RouteHandler{
		StreamTracker: &streamTracker,
	}
	mux := http.NewServeMux()
	SetupPaths(mux, advancedHandler)

	fs := http.FileServer(http.Dir("./static/"))
	handler := ServeWithoutHTMLExtension{fs: fs, staticFolder: "./static/"}
	mux.HandleFunc("/", handler.ServerHTTP)

	fmt.Printf("WebSocket server started on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func SetupPaths(mux *http.ServeMux, advancedHandler routehandlers.RouteHandler) {
	http.HandleFunc("/api/info", advancedHandler.InfoHandler)
	http.HandleFunc("/api/mkv", advancedHandler.MkvHandler)
	http.HandleFunc("/api/watch/mkv", advancedHandler.WatchMkv)
	http.HandleFunc("/api/backup", advancedHandler.BackupHandler)
	http.HandleFunc("POST /api/register", routehandlers.RegistrationHandler)
	http.HandleFunc("POST /api/eject", routehandlers.EjectHandler)
	http.HandleFunc("POST /api/insert", routehandlers.InsertDiscHandler)
}

func runInitialDiscLoadOnStartup() {
	//TODO: Set this value in config
	initalLoadReader, _, _ := makemkv.TriggerInitialInfoLoad(time.Minute * 2)
	stringChan := readStream(initalLoadReader)
	go func() {
		for {
			select {
			case newRead, ok := <-stringChan:
				if !ok {
					return
				}
				log.Println(newRead)
			}
		}
	}()
}

func readStream(reader io.Reader) <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
