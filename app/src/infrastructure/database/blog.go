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

func (ud *userDatabase) GetAllArticle() (*[]model.Article, error) {

	var articles []model.Article
	var err error
	return &articles, err
}

func (ud *userDatabase) GetOneArticle(id int) (*model.Article, error) {

	var article model.Article
	var err error
	return &article, err
}

// func (id *)
