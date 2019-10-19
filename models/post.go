package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Post model defines what is going to be stored in DB
type Post struct {
	ID         int
	Title      string
	Content    string
	InsertedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
	Active     bool
	User       *User `orm:"column(user_id);rel(fk)"`
}

// TableName returns normalized table name
func (p *Post) TableName() string {
	return "post"
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Post))
}
