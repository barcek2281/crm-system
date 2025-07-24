package service

import (
	"crmsystem/internal/config"
	"crmsystem/internal/dal"
	"crmsystem/internal/model"
	"crmsystem/internal/util"
	"database/sql"
	"errors"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyEmail      = errors.New("empty email")
	ErrPasswordSmall   = errors.New("password is less than 8 symbols")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInternal        = errors.New("internal problem")
	ErrEmailOrPassword = errors.New("login or password is not correct")
)

type Auth struct {
	authRepo *dal.User
	cnf      config.Config
	v        *validator.Validate
}

func NewAuthService(userrepo *dal.User, cnf config.Config) *Auth {
	return &Auth{
		authRepo: userrepo,
		cnf:      cnf,
		v:        validator.New(),
	}
}

func (a *Auth) Register(user model.RegisterRequest) (model.RegisterResponse, error) {
	if err := a.v.Struct(user); err != nil {
		return model.RegisterResponse{}, err
	}
	passHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		slog.Warn("eror with hash", "erro", err)
		return model.RegisterResponse{}, err
	}

	user.Password = string(passHashed)

	// bruh i hate this job
	user.Login = user.FirstName + strconv.Itoa(rand.Intn(1000))

	// TODO: find company name and job title

	id, err := a.authRepo.Register(user)
	if err != nil {
		slog.Warn("eror with db", "erro", err)
		return model.RegisterResponse{}, ErrInternal
	}

	token, err := util.CreateJWTsession([]byte(a.cnf.Srv.SecretJws),
		jwt.MapClaims{
			"user_id":  id,
			"is_admin": user.IsAdmin,
			"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
	if err != nil {
		slog.Warn("cannot create session", "error", err)
		return model.RegisterResponse{}, ErrInternal
	}

	res := model.RegisterResponse{
		Login:        user.FirstName + id[:5],
		ID:           id,
		SessionToken: token,
	}
	return res, nil
}

func (a *Auth) Login(user model.LoginRequest) (model.LoginResponse, error) {

	err := a.authRepo.Exist(user)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.LoginResponse{}, ErrEmailOrPassword
		}
		return model.LoginResponse{}, ErrInternal
	}

	user2, err := a.authRepo.GetByLogin(user.Login)
	if err != nil {
		slog.Info("error to get login", "error", err)
		return model.LoginResponse{}, ErrInternal
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user2.Password), []byte(user.Password)); err != nil {
		return model.LoginResponse{}, ErrEmailOrPassword
	}

	token, err := util.CreateJWTsession([]byte(a.cnf.Srv.SecretJws),
		jwt.MapClaims{
			"user_id":  user2.ID,
			"is_admin": user2.IsAdmin,
			"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
	if err != nil {
		slog.Warn("cannot create session", "error", err)
		return model.LoginResponse{}, ErrInternal
	}
	return model.LoginResponse{SessionToken: token, ID: user2.ID}, nil
}

func (a *Auth) IsAdmin(userId string) (bool, error) {
	user, err := a.authRepo.GetById(userId)
	if err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}
