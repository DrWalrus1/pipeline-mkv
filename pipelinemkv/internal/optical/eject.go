package optical

import (
	"log"
	"net/http"
	"os/exec"
)

func ejectHandler(w http.ResponseWriter, r *http.Request) {
	device := r.URL.Query().Get("device")
	log.Printf("Ejecting device: %s", device)

	responseStatus := ejectDevice(device)
	if responseStatus != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("Could not eject device: " + responseStatus.Error()))
		return
	}
	w.WriteHeader(200)
}

func ejectDevice(device string) error {
	// execute the bash command to eject the device
	cmd := exec.Command("eject", device)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
