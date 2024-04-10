package endpointsbaner

import (
	psql "banner/internal/repo/Db"
	servicebanner "banner/internal/service/serviceBanner"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (h *endpoint) Update(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		h.jsonError(w, http.StatusBadRequest, errors.New("no id"))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.s.Update(reqID, id, body)

	switch err := errors.Cause(err).(type) {
	case psql.ErrNotFound:
		h.jsonError(w, http.StatusNotFound, err)
		return
	case servicebanner.ErrIncorrectData:
		h.jsonError(w, http.StatusBadRequest, err)
		return
	case nil:
		w.WriteHeader(http.StatusOK)
	default:
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

}
