package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte("humbak")



func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Token")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			SendErrorJson(w, "no token header")
			return
		}

		token, err := jwt.Parse(tokenHeader, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				SendErrorJson(w, "error signing token")
				return nil, fmt.Errorf("error signing token")
			}
			return secret, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			SendErrorJson(w, "token error")
			return
		}

		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			SendErrorJson(w, "invalid authorization")
		}
	}
}
