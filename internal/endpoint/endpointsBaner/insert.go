package endpointsbaner

import (
	servicebanner "banner/internal/service/serviceBanner"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func (h *endpoint) Insert(w http.ResponseWriter, r *http.Request) {

	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	role := r.Context().Value("role").(bool)
	if !role {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

	id, err := h.s.Insert(reqID, body)
	switch err := errors.Cause(err).(type) {
	case servicebanner.ErrIncorrectData:
		h.jsonError(w, http.StatusBadRequest, err)
		return
	case nil:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]int{"banner_id": id})
	default:
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

}
