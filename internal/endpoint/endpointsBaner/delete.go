package endpointsbaner

import (
	psql "banner/internal/repo/Db"
	servicebanner "banner/internal/service/serviceBanner"
	"net/http"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

func (h *endpoint) Delete(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	vars := mux.Vars(r)
	t, ok := vars["id"]
	if !ok {
		h.jsonError(w, http.StatusBadRequest, errors.New("no id"))
	}

	err := h.s.Delete(reqID, t)
	switch err := errors.Cause(err).(type) {
	case psql.ErrNotFound:
		h.jsonError(w, http.StatusNotFound, err)
		return
	case servicebanner.ErrIncorrectData:
		h.jsonError(w, http.StatusBadRequest, err)
		return
	case nil:
		w.WriteHeader(http.StatusNoContent)
	default:
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}
}
