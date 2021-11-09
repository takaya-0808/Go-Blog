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

func (bd *blogDatabase) GetAllArticle() (*[]model.Article, error) {

	var articles []model.Article
	var err error
	rows, err := bd.Conn.Query("select * from　blogs")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		article := model.Article{}
		if err := rows.Scan(&article.ID, &article.Author, &article.Content, &article.Title, &article.CreatedAt, &article.UpdatedAt); err != nil {
			panic(err)
		}
		articles = append(articles, article)
	}

	return &articles, err
}

func (bd *blogDatabase) GetOneArticle(id int) (*model.Article, error) {

	var article model.Article
	var err error
	return &article, err
}

func (bd *blogDatabase) PostArticle(*model.Article) error {

	var err error
	return err
}
