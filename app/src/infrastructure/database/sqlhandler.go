package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewConnectionHandler() *SqlHandler {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbProtocol := "tcp(dbblog:3306)"
	dbName := "go_blog"
	dbOption := "?parseTime=true"

	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"/"+dbName+dbOption)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
