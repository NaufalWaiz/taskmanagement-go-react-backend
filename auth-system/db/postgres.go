package db

import (
    "database/sql"
    "log"
    "project/config"

    _ "github.com/jackc/pgx/v4/stdlib"
)

var Db *sql.DB

func InitPostgres() error {
    connStr := config.GetEnv("POSTGRESQL_URI", "user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
    var err error
    Db, err = sql.Open("pgx", connStr)
    if err != nil {
        return err
    }

    if err = Db.Ping(); err != nil {
        return err
    }

    log.Println("Successfully connected to PostgreSQL!")
    return nil
}