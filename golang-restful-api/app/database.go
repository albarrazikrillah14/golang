package app

import (
	"database/sql"
	"medomeckz/category-restful-api/helper"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func NewDB() *sql.DB {
	connStr := "user=medomeckz dbname=categories password=P_assword001 port=5432 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
