package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"servermakemkv/routehandlers"
	"strings"
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
	http.HandleFunc("/api/info", routehandlers.InfoHandler)
	http.HandleFunc("/api/mkv", routehandlers.MkvHandler)
	http.HandleFunc("/api/backup", routehandlers.BackupHandler)
	http.HandleFunc("POST /api/register", routehandlers.RegistrationHandler)

	fs := http.FileServer(http.Dir("./static/"))
	handler := ServeWithoutHTMLExtension{fs: fs, staticFolder: "./static/"}
	http.HandleFunc("/", handler.ServerHTTP)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
