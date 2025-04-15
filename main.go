package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"servermakemkv/outputs"
	"servermakemkv/parser"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development).  **SECURITY WARNING**:  In production, restrict this!
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case _ = <-ticker.C:
			titleInformation, err := parser.Parse("test")
			eventData := outputs.JsonWrapper{
				Type: titleInformation.GetTypeName(),
				Data: titleInformation,
			}

			jsonData, _ := json.Marshal(eventData)

			err = conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Println("write error:", err)
				return // Exit if we can't write (client likely disconnected)
			}
		}
	}
}

func main() {
	http.HandleFunc("/events", websocketHandler)

	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
