package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Initial_20191019_064444 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Initial_20191019_064444{}
	m.Created = "20191019_064444"

	migration.Register("Initial_20191019_064444", m)
}

// Run the migrations
func (m *Initial_20191019_064444) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(
		`CREATE TABLE IF NOT EXISTS user (
			id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
			name VARCHAR(255) NOT NULL DEFAULT '',
			password VARCHAR(255) NOT NULL DEFAULT '',
			inserted_at DATETIME NOT NULL,
			updated_at DATETIME NULL,
			active BOOLEAN NOT NULL
		);`,
	)
	m.SQL(
		`CREATE TABLE IF NOT EXISTS post (
			id INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
			title VARCHAR(255) NOT NULL DEFAULT '',
			content TEXT NOT NULL,
			inserted_at DATETIME NOT NULL,
			updated_at DATETIME NULL,
			active BOOLEAN NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id)
		);`,
	)
}

// Down Reverse the migrations
func (m *Initial_20191019_064444) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE user`)
	m.SQL(`DROP TABLE post`)
}
