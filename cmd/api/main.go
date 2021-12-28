package main

import (
	"backend/models"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	JWT struct {
		Secret string
	}
}

type AppStatus struct {
	Status       string `json:"status"`
	Environtment string `json:"environment"`
	Version      string `json:"version"`
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.Port, "Port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "Env", "Development", "Application environtment (development | production)")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "postgres://postgres:db@localhost/go_movies?sslmode=disable", "Postgres connection string")
	flag.StringVar(&cfg.JWT.Secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &Application{
		Config: cfg,
		Logger: logger,
		Models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
