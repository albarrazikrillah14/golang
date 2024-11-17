package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	connStr := "user=medomeckz dbname=belajar_golang password=P_assword001 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Hour)

	return db
}
