package golang_gorm

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id"`
	Password  string    `gorm:"column:password"`
	Name      Name      `gorm:"embedded"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime"`
	UpdatedAt time.Time `'gorm:"updated_at;autoCreateTime;autoUpdateTime"`
}

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
}

func (u *User) TableName() string {
	return "users"
}

type UserLog struct {
	ID        int       `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Action    string    `gorm:"column:action"`
	CreatedAt time.Time `gorm:"created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `'gorm:"updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}
