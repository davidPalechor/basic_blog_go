package utils

import (
	"github.com/astaxie/beego/orm"
)

// DBOrmer defines an interface with the basic CRUD operations
type DBOrmer interface {
	Insert(obj interface{}) (int64, error)
	One(queryable orm.QuerySeter, object interface{}, cols ...string) error
	GetQueryTable(model string) orm.QuerySeter
}

// DBWrapper defines an implementation of DBOrmer interface
type DBWrapper struct {
	db orm.Ormer
}

func NewDBWrapper() DBOrmer {
	return &DBWrapper{db: orm.NewOrm()}
}

func (d *DBWrapper) GetQueryTable(model string) orm.QuerySeter {
	return d.db.QueryTable(model)
}

// Insert defines the basi insert into a table statement
func (d *DBWrapper) Insert(obj interface{}) (int64, error) {
	return d.db.Insert(obj)
}

// One retrieve a unique registry from db
func (d *DBWrapper) One(queryable orm.QuerySeter, object interface{}, cols ...string) error {
	return queryable.One(object, cols...)
}
