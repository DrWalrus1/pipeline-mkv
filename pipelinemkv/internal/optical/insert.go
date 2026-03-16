package optical

import (
	"log"
	"net/http"
	"os/exec"
)

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

func InsertDevice(device string) error {
	// execute the bash command to insert the device
	cmd := exec.Command("eject", "-t", device)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
