package makemkv

import (
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
)

type registerHandler interface {
	RegisterMakeMkv(key string) error
}

func (h *MakeMkvRouteHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	err := h.CommandHandler.RegisterMakeMkv(key)
	if err != nil {
		switch err {
		case gomakemkv.ErrBadKey:
			w.WriteHeader(http.StatusBadRequest)
		case gomakemkv.ErrUnexpectedRegistrationError:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
