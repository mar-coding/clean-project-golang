package service

import (
	"context"
	"github.com/mar-coding/clean-project-golang/config"
	"github.com/mar-coding/clean-project-golang/internal/domain"
	blogPB "github.com/mar-coding/clean-project-golang/pb/proto-gen/services/blog/v1"
	converter "github.com/mar-coding/clean-project-golang/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b BlogService) CreateArticle(ctx context.Context, request *blogPB.CreateArticleRequest) (*blogPB.Article, error) {
	articleUseCase := domain.Bridge[domain.ArticleUseCase](config.ARTICLE_COLLECTION, b.useCases)

	articleModel, err := converter.ConvertProtoToModel[*domain.Article](request)
	if err != nil {
		return nil, err
	}

	article, err := articleUseCase.CreateArticle(ctx, articleModel)
	if err != nil {
		return nil, err
	}

	return converter.ConvertModelToProto[*blogPB.Article](article)
}

func (b BlogService) GetArticle(ctx context.Context, request *blogPB.GetArticleRequest) (*blogPB.Article, error) {
	articleUseCase := domain.Bridge[domain.ArticleUseCase](config.ARTICLE_COLLECTION, b.useCases)
	articleId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	article, err := articleUseCase.GetArticleById(ctx, articleId)
	if err != nil {
		return nil, err
	}

	articlePB, err := converter.ConvertModelToProto[*blogPB.Article](article)
	if err != nil {
		return nil, err
	}

	return articlePB, nil
}
