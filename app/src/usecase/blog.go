package usecase

import (
	"Go-Blog/app/src/domain/model"
	"Go-Blog/app/src/domain/repository"
)

type BlogUseCase interface {
	Index() (*[]model.Article, error)
}

type blogUseCases struct {
	blogRepository repository.BlogRepository
}

func NewBlogUseCase(br repository.BlogRepository) BlogUseCase {
	return &blogUseCases{
		blogRepository: br,
	}
}

func (bu blogUseCases) Index() (*[]model.Article, error) {

	articles, err := bu.blogRepository.GetAllArticle()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
