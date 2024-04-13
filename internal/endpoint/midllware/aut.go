package midllware

import (
	"banner/internal/config"
	"banner/internal/model"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var Users model.Users = []model.User{
	{
		Username: "admin",
		Root:     true,
	},
	{
		Username: "user",
		Root:     false,
	}}

func (m *midllware) Aut(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := r.Context().Value("reqID").(string)
		if !ok {
			reqID = ""
		}

		token := r.Header.Get("token")
		if token != "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &model.Claims{}

		cfg := config.TokenCFG()

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.SecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, user := range Users {
			if user.Username == claims.Username {
				ctx := context.WithValue(r.Context(), "reqID", reqID)
				ctx = context.WithValue(ctx, "role", user.Root)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

	})
}
