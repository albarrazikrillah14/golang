package app

import (
	"database/sql"
	"medomeckz/category-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	connStr := "user=medomeckz dbname=categories  password=P_assword001 port=5432 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(1 * time.Minute)

	return db
}
