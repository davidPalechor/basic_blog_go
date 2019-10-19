package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

const UserTableName = "user"

// User model defines structure of users table
type User struct {
	ID         int       `orm:"column(id)"`
	Name       string    `orm:"column(name)"`
	Email      string    `orm:"column(email);unique"`
	Password   string    `orm:"column(password)"`
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
	Active     bool      `orm:"column(active);default(1)"`
}

// TableName returns normalized table name
func (p *User) TableName() string {
	return UserTableName
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}
