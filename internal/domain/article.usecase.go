package domain

import "context"

type ArticleUseCase interface {
	Bridger
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
}
