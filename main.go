package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"servermakemkv/commands"
	"strconv"

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
	source := r.URL.Query().Get("source")
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
	source := r.URL.Query().Get("source")
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

func backupHandler(w http.ResponseWriter, r *http.Request) {
	decrypt, err := strconv.ParseBool(r.URL.Query().Get("decrypt"))
	if err != nil {
		r.Response.StatusCode = 400
		return
	}
	source := r.URL.Query().Get("source")
	destination := r.URL.Query().Get("destination")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	updates := make(chan []byte)
	// TODO: add test for cancelChannel
	cancelChannel := make(chan bool)
	// TODO: add error handling
	go commands.BackupDisk(decrypt, source, destination, updates, cancelChannel)

	go func() {
		defer close(updates)
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
			}
			if string(p) == "cancel" {
				cancelChannel <- true
				close(cancelChannel)
			}
		}
	}()
	for update := range updates {
		err = conn.WriteMessage(websocket.TextMessage, update)
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
	}
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	// TODO: add error handling
	responseStatus := commands.RegisterMkvKey(key)
	w.WriteHeader(responseStatus)
}

func main() {
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/mkv", mkvHandler)
	http.HandleFunc("/backup", backupHandler)
	http.HandleFunc("POST /register", registrationHandler)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
