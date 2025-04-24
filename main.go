package main

import (
	"fmt"
	"log"
	"net/http"
	"servermakemkv/commands"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development).  **SECURITY WARNING**:  In production, restrict this!
	},
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	source := r.PathValue("source")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	updates := make(chan []byte)
	// TODO: add error handling
	go commands.GetInfo(nil, source, updates)
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
}

func mkvHandler(w http.ResponseWriter, r *http.Request) {
	source := r.PathValue("source")
	title := r.URL.Query().Get("title")
	destination := r.URL.Query().Get("destination")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	updates := make(chan []byte)
	// TODO: add error handling
	go commands.SaveMkv(source, title, destination, updates)
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
}

func main() {
	http.HandleFunc("/info/{source}", infoHandler)
	http.HandleFunc("/mkv/{source}", mkvHandler)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
