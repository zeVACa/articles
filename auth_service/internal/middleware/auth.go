package mymiddleware

import (
	"articles/pgk/myjson"
	"articles/pgk/myjwt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Привет из мидлвары")

		type Body struct {
			Token string `json:"token"`
		}
		var body Body
		json.NewDecoder(r.Body).Decode(&body)
		log.Println("Токен", body)

		_, err := myjwt.ValidateToken(body.Token)
		if err != nil {
			myjson.SendError(w, err.Error(), "Невалидный токен", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}
