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

	// TODO: add error handling
	reader, cancel, err := commands.TriggerDiskInfo(source)
	if err != nil {
		log.Printf("Could not trigger get disk info: %v", err)
		cancel()
		err = conn.WriteMessage(websocket.TextMessage, fmt.Appendf(nil, "Could not trigger get disk info: %v", err))
		if err != nil {
			log.Println("write error:", err)
			return // Exit if we can't write (client likely disconnected)
		}
		return
	}
	updates := commands.WatchInfoLogs(reader)
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				cancel()
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
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

	// TODO: add error handling
	reader, cancel, _ := commands.TriggerSaveMkv(source, title, destination)
	updates := commands.WatchSaveMkvLogs(reader)
	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				cancel()
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
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

	// TODO: add error handling
	reader, cancel, _ := commands.TriggerDiskBackup(decrypt, source, destination)
	updates := commands.WatchBackupLogs(reader)

	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if string(p) == "cancel" {
				cancel()
				return
			}
			if err != nil {
				log.Println("read error:", err)
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) || err == io.EOF {
					return
				}
				return
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
	http.HandleFunc("/api/info", infoHandler)
	http.HandleFunc("/api/mkv", mkvHandler)
	http.HandleFunc("/api/backup", backupHandler)
	http.HandleFunc("POST /api/register", registrationHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
