package server

import (
	"crmsystem/internal/config"
	"crmsystem/internal/dal"
	"crmsystem/internal/handler"
	"crmsystem/internal/service"
	"fmt"
	"net/http"
)

type Server struct {
	authHandler    *handler.Auth
	jobHandler     *handler.Job
	companyHandler *handler.Company
	cnf            config.Config
	mux            *http.ServeMux
}

func NewServer(cnf config.Config) (*Server, error) {
	store, err := dal.NewStore(cnf)
	if err != nil {
		return nil, err
	}

	userRepo := store.User()
	jobRepo := store.Job()
	comRepo := store.Company()

	authService := service.NewAuthService(userRepo, cnf)
	jobService := service.NewJobService(jobRepo)
	companyService := service.NewCompanyService(comRepo)

	authHandler := handler.NewAuthHandler(authService)
	jobHandler := handler.NewJobHandler(jobService)
	company := handler.NewCompanyHandler(companyService)

	return &Server{
		companyHandler: company,
		jobHandler:     jobHandler,
		authHandler:    authHandler,
		cnf:            cnf,
		mux:            http.NewServeMux(),
	}, nil
}

func (s *Server) Init() {
	s.mux.HandleFunc("POST /create/company", s.jobHandler.CreateJob())
	s.mux.HandleFunc("POST /create/job", s.jobHandler.CreateJob())
	s.mux.HandleFunc("POST /register", s.authHandler.RegisterUser())
	s.mux.HandleFunc("POST /login", s.authHandler.Login())
}

func (s *Server) Run() error {
	s.Init()

	return http.ListenAndServe(fmt.Sprintf(":%d", s.cnf.Srv.Port), s.mux)
}
