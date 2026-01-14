package mymiddleware

import (
	"articles/pkg/jsonPkg"
	"articles/pkg/jwtPkg"
	"encoding/json"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Token string `json:"token"`
		}
		var body Body
		json.NewDecoder(r.Body).Decode(&body)

		_, err := jwtPkg.ValidateToken(body.Token)
		if err != nil {
			jsonPkg.SendError(w, err.Error(), "Невалидный токен", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
