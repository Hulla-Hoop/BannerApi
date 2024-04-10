package endpointsbaner

import (
	psql "banner/internal/repo/Db"
	servicebanner "banner/internal/service/serviceBanner"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (h *endpoint) GetBanner(w http.ResponseWriter, r *http.Request) {

	reqID, ok := r.Context().Value("reqID").(string)
	if !ok {
		reqID = ""
	}

	query := r.URL.Query()

	tagID, ok := query["tag_id"]
	if !ok {
		h.jsonError(w, http.StatusBadRequest, errors.New("no tag_id"))
		return
	}
	featureID, ok := query["feature_id"]
	if !ok {
		h.jsonError(w, http.StatusBadRequest, errors.New("no feature_id"))
		return
	}
	last, ok := query["use_last_revision"]
	if !ok {
		last[0] = "false"
	}

	b, err := h.s.GetOne(reqID, tagID[0], featureID[0], last[0])
	switch err := errors.Cause(err).(type) {
	case psql.ErrNotFound:
		h.jsonError(w, http.StatusNotFound, err)
		return
	case servicebanner.ErrIncorrectData:
		h.jsonError(w, http.StatusBadRequest, err)
		return
	case nil:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(b)
	default:
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

}
