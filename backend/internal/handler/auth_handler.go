package handler

import (
	"crmsystem/internal/model"
	"crmsystem/internal/service"
	"crmsystem/internal/util"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Auth struct {
	authService *service.Auth
}

func NewAuthHandler(authService *service.Auth) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (a *Auth) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.RegisterRequest{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "json encoding problem", http.StatusBadRequest)
			slog.Warn("cannot read json", "error", err)
			return
		}

		res, err := a.authService.Register(user)
		if err != nil {
			http.Error(w, "cannot register", http.StatusBadRequest)
			slog.Warn("error to register", "error", err)
			return
		}

		util.ResponseJSON(w, http.StatusOK, res)
	}
}

func (a *Auth) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		login := model.LoginRequest{}

		if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
			http.Error(w, "json encoding problem", http.StatusBadRequest)
			slog.Warn("cannot read json", "error", err)
			return
		}

		res, err := a.authService.Login(login)

		if err != nil {
			if err == service.ErrEmailOrPassword {
				http.Error(w, err.Error(), http.StatusBadRequest)
			} else {
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}

		util.ResponseJSON(w, http.StatusOK, res)

	}
}
