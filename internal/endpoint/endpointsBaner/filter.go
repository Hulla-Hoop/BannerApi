package endpointsbaner

import (
	"encoding/json"
	"net/http"
)

func (h *endpoint) Filter(w http.ResponseWriter, r *http.Request) {
	reqID, ok := r.Context().Value("reqID").(string)

	if !ok {
		reqID = ""
	}
	//при отсутствии параметров возвращаем все баннеры
	feature := r.URL.Query().Get("feature_id")

	tag := r.URL.Query().Get("tag_id")

	limit := r.URL.Query().Get("limit")

	offset := r.URL.Query().Get("offset")

	banners, err := h.s.Filter(reqID, tag, feature, limit, offset)
	if err != nil {
		h.jsonError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banners)
}
