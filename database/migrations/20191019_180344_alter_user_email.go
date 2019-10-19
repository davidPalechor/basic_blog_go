package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AlterUserEmail_20191019_180344 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AlterUserEmail_20191019_180344{}
	m.Created = "20191019_180344"

	migration.Register("AlterUserEmail_20191019_180344", m)
}

// Run the migrations
func (m *AlterUserEmail_20191019_180344) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE user ADD COLUMN email VARCHAR(100) NOT NULL UNIQUE;")

}

// Reverse the migrations
func (m *AlterUserEmail_20191019_180344) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE user DROP COLUMN email;")
}
