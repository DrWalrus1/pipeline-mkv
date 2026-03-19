package config

import (
	"encoding/json"
	"io"
	"net/http"
)

func SetupConfigApiPaths(mux *http.ServeMux, c *Config) {
	mux.HandleFunc("GET /api/config", makeGetConfigHandler(c))
	mux.HandleFunc("PUT /api/config", makeUpdateConfigHandler(c))
}

func makeGetConfigHandler(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(data))
		w.WriteHeader(http.StatusOK)
	}

}

func makeUpdateConfigHandler(c *Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var newConf Config
		err = json.Unmarshal(body, &newConf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		c = &newConf
		w.WriteHeader(http.StatusAccepted)
	}
}
