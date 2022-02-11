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

	rows, err := bd.Conn.Query("select * from blogs")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var articles []model.Article
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
	err := bd.Conn.QueryRow("select * from blogs where id = ?LIMIT 1", id).Scan(&article)
	if err != nil {
		return nil, err
	}
	return &article, err
}

func (bd *blogDatabase) TitleShow() (*[]model.TitlesShow, error) {

	var titles []model.TitlesShow
	rows, err := bd.Conn.Query("select * from titles")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		title := model.TitlesShow{}
		if err := rows.Scan(&title.ID, &title.Title, &title.URL); err != nil {
			panic(err)
		}
		titles = append(titles, title)
	}
	return &titles, err
}

func (bd *blogDatabase) PostArticle(article *model.CreateArticle) error {

	var err error
	_, err = bd.Conn.Exec("insert into blogs (author, title, context) VALUES (?, ?, ?)", article.Author, article.Title, article.Content)
	if err != nil {
		return err
	}
	return err
}
