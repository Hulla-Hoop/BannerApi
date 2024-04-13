package midllware

import (
	"banner/internal/config"
	"banner/internal/model"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var Users model.Users = []model.User{
	// admin eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIn0.v4A2q3xg9-zFWjP_CTV2HQuteszG7Mx08GUOLUIfOnG1a2P9c2ZU1FKKRiVEpJVMMZvCb4JjlPTNrkzIy1tCbA
	{
		Username: "admin",
		Root:     true,
	},
	// user eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIifQ.IQohLkYqxzY9A6ent4MGs1NyBNcSQyiyAd5ZG_c39CEHbuKwOuMNXhMO5dg01rB9CSV5R7MchcaZHDYZs_k7Bg
	{
		Username: "user",
		Root:     false,
	}}

// chel(для теста доступа) eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNoZWwifQ.2ZqHPnzoswMpS4WuvSFQdO97KoS0GYYD71kbY9dPTKl4RXFs1TtDYWlOBJn7liCE1eeozVn3Fgew-RxI7ZA-pQ

func (m *midllware) Aut(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := r.Context().Value("reqID").(string)
		if !ok {
			reqID = ""
		}

		tok := r.Header

		token := tok["Token"]
		m.L.WithField("middleware", reqID).Debug("Header --- ", token)
		if token[0] == "" {
			m.L.WithField("middleware", reqID).Debug("token", token)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &model.Claims{}

		cfg := config.TokenCFG()

		tkn, err := jwt.ParseWithClaims(token[0], claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.SecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				m.L.WithField("middleware", reqID).Error(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			m.L.WithField("middleware", reqID).Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			m.L.WithField("middleware", reqID).Debug("invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		for _, user := range Users {
			if user.Username == claims.Username {
				m.L.WithField("middleware", reqID).Debug("user --- ", user.Username, " -- Name in claims --", claims.Username)
				ctx := context.WithValue(r.Context(), "reqID", reqID)
				ctx = context.WithValue(ctx, "role", user.Root)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	})
}
