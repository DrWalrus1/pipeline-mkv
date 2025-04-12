package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func ServerSideEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := r.Context().Done()

	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-clientGone:
			fmt.Println("Client disconnected")
			return
		case <-t.C:
			// Send an event to the client
			// Here we send only the "data" field, but there are few others
			eventTest := struct {
				Hello string `json:"hello"`
			}{
				Hello: "hello there",
			}
			stringifiedEvent, _ := json.Marshal(eventTest)
			_, err := fmt.Fprintf(w, "%s\n\n", string(stringifiedEvent))
			if err != nil {
				return
			}
			err = rc.Flush()
			if err != nil {
				return
			}
		}
	}
}

func main() {
	http.HandleFunc("/events", ServerSideEventHandler)
	fmt.Println("server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
