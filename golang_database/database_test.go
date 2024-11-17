package golang_database

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	connStr := "user=medomeckz dbname=belajar_golang password=P_assword001 host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Hour)
	defer db.Close()
}
