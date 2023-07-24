package usecase

import (
	"context"
	"fmt"
	"github.com/mar-coding/clean-project-golang/internal/domain"
)

type ArticleUseCase struct {
	//errHandler   errHandler.Handler
	//logger       zapper.Zapper
	//user         userPB.UserServiceClient
	repositories map[string]domain.Bridger
}

func (a *ArticleUseCase) CreateArticle(ctx context.Context, article *domain.Article) (*domain.Article, error) {
	fmt.Print("to be implemented")
	return nil, nil
}
