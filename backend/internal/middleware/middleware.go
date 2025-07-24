package middleware

import (
	"context"
	"crmsystem/internal/config"
	"crmsystem/internal/service"
	"crmsystem/internal/util"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type sessionKey string

var session sessionKey = "user_id"
var ErrUnauthorized = errors.New("not auth")

type Mid struct {
	auth *service.Auth
	cnf  config.Config
}

func NewMid(auth *service.Auth, cnf config.Config) *Mid {
	return &Mid{
		auth: auth,
		cnf:  cnf,
	}
}

func (m *Mid) Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			util.ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}
		t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(m.cnf.Srv.SecretJws), nil
		})

		if err != nil {
			util.ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			util.ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		userID, ok := claims[string(session)].(string)
		if !ok {
			util.ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}

		ok, err = m.auth.IsAdmin(userID)
		if err != nil {
			util.ResponseError(w, http.StatusInternalServerError, nil)
			return
		}

		if !ok {
			util.ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), session, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
