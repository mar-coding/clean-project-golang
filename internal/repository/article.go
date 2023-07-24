package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/mar-coding/clean-project-golang/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleRepository struct {
	//errHandler errHandler.Handler
	//logger     zapper.Zapper
	collection *mongo.Collection
}

func (a *ArticleRepository) CreateArticle(ctx context.Context, article *domain.Article) error {
	if _, err := a.collection.InsertOne(ctx, article); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("article already exists")
		}
		return errors.New(fmt.Sprintf("failed to create article. got error %s", err.Error()))
	}
	return nil
}
func (a *ArticleRepository) ListArticles(ctx context.Context) ([]*domain.Article, error) {
	fmt.Print("to be implemented")
	return nil, nil
}
func (a *ArticleRepository) GetArticle(ctx context.Context, articleId primitive.ObjectID) (*domain.Article, error) {
	fmt.Print("to be implemented")
	return nil, nil
}
func (a *ArticleRepository) UpdateArticle(ctx context.Context, articleId primitive.ObjectID, fields *domain.UpdateArticleFields) error {
	fmt.Print("to be implemented")
	return nil
}
func (a *ArticleRepository) DeleteArticle(ctx context.Context, articleId primitive.ObjectID) error {
	fmt.Print("to be implemented")
	return nil
}
