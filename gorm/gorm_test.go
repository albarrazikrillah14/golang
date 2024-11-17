package golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=medomeckz password=P_assword001 dbname=belajar_golang_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func TestConnection(t *testing.T) {
	db := OpenConnection()

	assert.NotNil(t, db)
}

var db *gorm.DB = OpenConnection()

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO sample(id, name) VALUES($1, $2)", "1", "Zikri").Error
	assert.Nil(t, err)
}
