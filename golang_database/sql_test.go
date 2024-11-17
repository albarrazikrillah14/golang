package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('joko', 'Wi')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT * FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name sql.NullString

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %s, Name: %s\n", id, name)
	}
	defer rows.Close()
}
