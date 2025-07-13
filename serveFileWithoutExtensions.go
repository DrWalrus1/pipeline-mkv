package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
