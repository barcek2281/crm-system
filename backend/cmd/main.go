package main

import (
	"crmsystem/internal/config"
	"crmsystem/internal/server"
	"log"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	cnf := config.Config{}

	d, err := os.ReadFile("env.toml")
	if err != nil {
		log.Fatalf("env.toml error: %v", err)
	}
	if err := toml.Unmarshal(d, &cnf); err != nil {
		log.Fatalf("error with toml file: %v", err)
	}

	s, err := server.NewServer(cnf)
	if err != nil {
		log.Fatalf("cannot init server: %v", err)
	}
	
	slog.Info("server listing", "port", cnf.Srv.Port)
	if err := s.Run(); err != nil {
		slog.Error("server problem", "error", err)
	}
}
