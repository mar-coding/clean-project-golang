package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleUseCase interface {
	Bridger
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
	ListArticles(ctx context.Context) ([]*Article, error)
	GetArticleById(ctx context.Context, articleId primitive.ObjectID) (*Article, error)
	GetArticleByTitle(ctx context.Context, articleId primitive.ObjectID) (*Article, error)
	UpdateArticle(ctx context.Context, articleId primitive.ObjectID, article *Article) error
	DeleteArticle(ctx context.Context, articleId primitive.ObjectID) error
}
