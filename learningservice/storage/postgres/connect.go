package postgres

import (
	"database/sql"
	"fmt"
	"learn_service/config"
	"log/slog"


	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db  *sql.DB
	log *slog.Logger
}

func NewLearnStorage(db *sql.DB, log *slog.Logger) *PostgresStorage {
	return &PostgresStorage{
		db:  db,
		log: log,
	}
}

func ConnectionDb() (*sql.DB, error) {
	conf := config.Load()
	conDb := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		conf.Postgres.DB_HOST, conf.Postgres.DB_PORT, conf.Postgres.DB_USER, conf.Postgres.DB_NAME, conf.Postgres.DB_PASSWORD)
	db, err := sql.Open("postgres", conDb)
	fmt.Println("connection", conDb)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}


