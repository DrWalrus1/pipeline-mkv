package makemkv

import (
	"net/http"

	"github.com/DrWalrus1/gomakemkv"
)

type registerHandler interface {
	RegisterMakeMkv(key string) error
}

func GetRegisterHandler(h registerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		err := h.RegisterMakeMkv(key)
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
}
