package usecase

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
	"errors"
)

type BlogUseCase interface {
	Index() (*[]model.Article, error)
	TitleOneShow(id int) (*model.TitlesShow, error)
	TitleShow() (*[]model.TitlesShow, error)
	GetArticle(id int) (*model.Article, error)
	Create(article *model.CreateArticle) error
	Update(article *model.CreateArticle) error
}

type blogUseCases struct {
	blogRepository repository.BlogRepository
}

func NewBlogUseCase(br repository.BlogRepository) BlogUseCase {
	return &blogUseCases{
		blogRepository: br,
	}
}

// GET function
func (bu blogUseCases) Index() (*[]model.Article, error) {

	articles, err := bu.blogRepository.GetAllArticle()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (bu blogUseCases) TitleShow() (*[]model.TitlesShow, error) {

	titles, err := bu.blogRepository.TitleShow()
	if err != nil {
		return nil, errors.New("No title")
	}
	print(titles)
	return titles, nil
}

func (bu blogUseCases) GetArticle(id int) (*model.Article, error) {

	article, err := bu.blogRepository.GetOneArticle(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (bu blogUseCases) TitleOneShow(id int) (*model.TitlesShow, error) {
	return nil, nil
}

// POST function
func (bu blogUseCases) Create(article *model.CreateArticle) error {

	err := bu.blogRepository.PostArticle(article)
	if err != nil {
		return err
	}
	return nil
}

// PUT function
func (bu blogUseCases) Update(article *model.CreateArticle) error {
	return nil
}
