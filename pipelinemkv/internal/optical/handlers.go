package optical

import (
	"log"
	"net/http"
)

func EjectHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	log.Printf("Ejecting device: %s", device)

	responseStatus := EjectDevice(device)
	if responseStatus != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not eject device: " + responseStatus.Error()))
		return
	}
	w.WriteHeader(200)
}

func InsertDiscHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	log.Printf("Inserting device: %s", device)

	responseStatus := InsertDevice(device)
	if responseStatus != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not insert device: " + responseStatus.Error()))
		return
	}
	w.WriteHeader(200)
}
