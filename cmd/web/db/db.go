package db

import (
	"database/sql"
	"os"

	"github.com/goququ/snippetbox/cmd/web/logger"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func Open() (*sql.DB, error) {
	if db != nil {
		logger.Info.Println("DB was initialized earlier")
		return db, nil
	}

	logger.Info.Println("Opening database...")
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info.Println("Database opened successful")

	return db, nil
}
