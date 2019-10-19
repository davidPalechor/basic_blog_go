package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User model defines structure of users table
type User struct {
	ID         int
	Name       string
	Password   string
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
	Active     bool
}

// TableName returns normalized table name
func (p *User) TableName() string {
	return "user"
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}
