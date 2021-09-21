package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	dsn  string
}

type application struct {
	config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.dsn, "dsn", "sample.db", "sqlite data source name")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	db, err := openDB(cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	app := &application{
		config: cfg,
		logger: logger,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logger.Printf("starting server on %s", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
