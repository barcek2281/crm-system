package service

import (
	"crmsystem/internal/config"
	"crmsystem/internal/dal"
	"crmsystem/internal/model"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var ErrEmptyEmail = errors.New("empty email")

type Auth struct {
	authRepo *dal.User
	cnf      config.Config
}

func NewAuthService(userrepo *dal.User, cnf config.Config) *Auth {
	return &Auth{
		authRepo: userrepo,
	}
}

func (a *Auth) Register(user model.RegisterRequest) (model.RegisterResponse, error) {
	if user.Email == "" {
		return model.RegisterResponse{}, ErrEmptyEmail
	}

	passHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		slog.Error("eror with hash", "erro", err)
		return model.RegisterResponse{}, err
	}
	user.Password = string(passHashed)
	if err := a.authRepo.Register(user); err != nil {
		slog.Error("eror with db", "erro", err)
		return model.RegisterResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		})

	tokenString, err := token.SignedString([]byte(a.cnf.Srv.SecretJws))
	if err != nil {
		slog.Error("eror with jwt", "error", err)
		return model.RegisterResponse{}, err
	}
	res := model.RegisterResponse{
		SessionToken: tokenString,
	}
	return res, err
}
