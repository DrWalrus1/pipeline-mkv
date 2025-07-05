package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"pipelinemkv/routehandlers"
	"strings"
	"time"

	"github.com/DrWalrus1/gomakemkv"
	"github.com/DrWalrus1/gomakemkv/commands"
)

type ServeWithoutHTMLExtension struct {
	fs           http.Handler
	staticFolder string
}

func (s ServeWithoutHTMLExtension) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "/") && !strings.Contains(filepath.Base(r.URL.Path), ".") {
		// If the path doesn't end in a slash and doesn't have an extension,
		// assume it's an HTML file and try appending ".html".
		newPath := r.URL.Path + ".html"
		_, err := os.Stat(filepath.Join(s.staticFolder, newPath)) // Check if the file exists
		if err == nil {
			r.URL.Path = newPath // Rewrite the URL internally
		} else {
			fmt.Printf("Couldn't find path %v", err)
		}
	}
	s.fs.ServeHTTP(w, r)
}

func main() {
	runInitialDiscLoadOnStartup()
	streamTracker := gomakemkv.NewStreamTracker()
	advancedHandler := routehandlers.RouteHandler{
		StreamTracker: &streamTracker,
	}
	http.HandleFunc("/api/info", advancedHandler.InfoHandler)
	http.HandleFunc("/api/mkv", advancedHandler.MkvHandler)
	http.HandleFunc("/api/watch/mkv", advancedHandler.WatchMkv)
	http.HandleFunc("/api/backup", advancedHandler.BackupHandler)
	http.HandleFunc("POST /api/register", routehandlers.RegistrationHandler)
	http.HandleFunc("POST /api/eject", routehandlers.EjectHandler)
	http.HandleFunc("POST /api/insert", routehandlers.InsertDiscHandler)

	fs := http.FileServer(http.Dir("./static/"))
	handler := ServeWithoutHTMLExtension{fs: fs, staticFolder: "./static/"}
	http.HandleFunc("/", handler.ServerHTTP)

	port := ":9090"
	fmt.Printf("WebSocket server started on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func runInitialDiscLoadOnStartup() {
	//TODO: Set this value in config
	initalLoadReader, _, _ := commands.TriggerInitialInfoLoad(time.Minute * 2)
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
