package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/w0/tinyIPAM/internal/database"
	_ "modernc.org/sqlite"
)

type application struct {
	logger   *slog.Logger
	database *database.Queries
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "file:ipam.db?&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)", "database connection string")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger:   logger,
		database: database.New(db),
	}

	logger.Info("starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
