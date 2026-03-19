package optical

import "net/http"

func SetupOpticalApiPaths(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/eject", ejectHandler)
	mux.HandleFunc("POST /api/insert", insertDiscHandler)
}
