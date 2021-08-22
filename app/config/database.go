package config

import (
	"database/sql"
)

var (
	db  *sql.DB
	err error
)

func Connection() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbProtocol := "tcp(dbblog:3306)"
	dbName := "go_blog"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
