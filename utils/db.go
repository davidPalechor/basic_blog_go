package utils

import (
	"github.com/astaxie/beego/orm"
)

// DBOrmer defines an interface with the basic CRUD operations
type DBOrmer interface {
	Insert(obj interface{}) (int64, error)
}

// DBWrapper defines an implementation of DBOrmer interface
type DBWrapper struct {
	db orm.Ormer
}

func NewDBWrapper() DBOrmer {
	return &DBWrapper{db: orm.NewOrm()}
}

// Insert defines the basi insert into a table statement
func (d *DBWrapper) Insert(obj interface{}) (int64, error) {
	return d.db.Insert(obj)
}
