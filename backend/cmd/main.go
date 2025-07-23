package main

import (
	"crmsystem/internal/config"
	"fmt"
	"log"
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

	fmt.Printf("%+v", cnf)
}
