package golang_gorm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dsn := "host=localhost user=medomeckz password=P_assword001 dbname=golang_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExcecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO sample(id, name) VALUES($1, $2)", "1", "Albarra").Error
	assert.NotNil(t, err)
	err = db.Exec("INSERT INTO sample(id, name) VALUES($1, $2)", "2", "Budi").Error
	assert.NotNil(t, err)
	err = db.Exec("INSERT INTO sample(id, name) VALUES($1, $2)", "3", "Joko").Error
	assert.NotNil(t, err)
	err = db.Exec("INSERT INTO sample(id, name) VALUES($1, $2)", "4", "Rully").Error
	assert.NotNil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("SELECT id, name FROM sample WHERE id = $1", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Albarra", sample.Name)

	samples := []Sample{}

	err = db.Raw("SELECT id, name FROM sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := &User{
		ID:       "1",
		Password: "P_assword001",
		Name: Name{
			FirstName:  "Albarra",
			MiddleName: "",
			LastName:   "Zikrillah",
		},
	}

	response := db.Create(user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User
	for i := 2; i <= 10; i++ {
		users = append(users, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}

	response := db.Create(users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(9), response.RowsAffected)
}

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "11",
			Password: "rahasia",
			Name: Name{
				FirstName: "User 11",
			},
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "12",
			Password: "rahasia",
			Name: Name{
				FirstName: "User 12",
			}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "13",
			Password: "rahasia",
			Name: Name{
				FirstName: "User 12",
			},
		}).Error

		if err != nil {
			return err
		}
		return nil
	})

	assert.Nil(t, err)
}

func TestQueryCondition(t *testing.T) {
	var users []User
	err := db.Where("first_name like $1", "%User%").Where("password = $2", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	fmt.Println(users)
	fmt.Println(len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	err := db.Where("first_name like $1", "%User%").Or("password = $2", "rahasia").Find(&users).Error

	assert.Nil(t, err)
	assert.Equal(t, 12, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User
	err := db.Not("password = $1", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "User 5",
		},
	}

	var users []User
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	userCondition := map[string]interface{}{
		"middle_name": "",
	}

	var users []User
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	users := []UserResponse{}
	err := db.Model(&User{}).Find(&users).Error
	assert.Nil(t, err)

	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", "1")
	assert.Nil(t, result.Error)

	user.Name.FirstName = "Sumbul"
	user.Password = "secret"

	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestUpdates(t *testing.T) {
	user := User{
		Name: Name{
			FirstName: "Medomeckz",
			LastName:  "Morro",
		},
	}
	err := db.Where("id = ?", "1").Updates(&user).Error
	assert.Nil(t, err)

	err = db.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Model(&User{}).Updates(map[string]interface{}{
		"first_name": "Kiji",
		"last_name":  "",
	}).Error

	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		err := db.Create(&userLog).Error
		assert.Nil(t, err)

		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}

	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)

	assert.Equal(t, 0, len(todos))

}
