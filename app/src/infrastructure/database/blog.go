package database

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"database/sql"
)

type blogDatabase struct {
	Conn *sql.DB
}

func NewBlogDabase(conn *sql.DB) repository.BlogRepository {

	return &blogDatabase{
		Conn: conn,
	}
}

func (ud *userDatabase) GetByID(*model.UserModel) int64 {

	var id int64
	return id
}
