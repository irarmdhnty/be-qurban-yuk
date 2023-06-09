package midleware

import (
	"context"
	"encoding/json"
	"net/http"
	"qurban-yuk/dto"
	"qurban-yuk/pkg/jwt"
	"strings"
)

type Result struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrResult{Status: "Failed", Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}
		token = strings.Split(token, " ")[1]
		claims, err := jwtauth.DecodeToken(token)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := Result{Status: "Failed", Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}
		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
