package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateArticleFields struct {
	Title   string
	Content string
}

type ArticleRepository interface {
	Bridger
	CreateArticle(ctx context.Context, article *Article) error
	ListArticles(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, articleId primitive.ObjectID) (*Article, error)
	UpdateArticle(ctx context.Context, articleId primitive.ObjectID, fields *UpdateArticleFields) error
	DeleteArticle(ctx context.Context, articleId primitive.ObjectID) error
}
