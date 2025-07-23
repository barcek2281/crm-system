package util

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func ResponseJSON(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if res != nil {
		json.NewEncoder(w).Encode(res)
	}
}

func ResponseError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"err": err.Error()})
	}
}

func CreateJWTsession(secret []byte, clm jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		clm,
	)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		slog.Error("eror with jwt", "error", err)
		return "", err
	}
	return tokenString, nil
}
