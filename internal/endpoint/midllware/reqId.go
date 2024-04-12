package midllware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Добавляет reqID в контекст для удобного отслеживания логов
func (m *midllware) ReqID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), "reqID", reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
