package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"marcode.io/phone-numbers/pkg/data"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	port int
	dsn  string
}

type application struct {
	config
	logger *log.Logger
	models data.Models
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
		models: data.NewModels(db),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logger.Printf("starting server on %s", server.Addr)

	go func() {
		err = server.ListenAndServe()
		logger.Fatal(err)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	// Graceful Shutdown for the server
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = server.Shutdown(tc)
	if err != nil {
		logger.Println(err)
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
