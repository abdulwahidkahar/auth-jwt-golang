package main

import (
	"log"
	"os"
	"auth-jwt-golang/internal/config"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config.LoadConfig()
	cfg := config.Config

	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	m, err := migrate.New(
		"file://db/migrations",
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration up done")
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migration down done")
	} else {
		log.Println("Unknown command. Use 'up' or 'down'")
	}
}
