package service

import (
	"github.com/mar-coding/clean-project-golang/config"
	"github.com/mar-coding/clean-project-golang/internal/domain"
	blogPB "github.com/mar-coding/clean-project-golang/pb/proto-gen/services/blog/v1"
	"google.golang.org/grpc"
)

type BlogService struct {
	blogPB.UnsafeBlogServiceServer
	useCases map[string]domain.Bridger
	//errHandler errHandler.Handler
	//logger     zapper.Zapper
}

func NewBlog(server *grpc.Server, useCases ...domain.Bridger) {
	blog := &BlogService{
		useCases: make(map[string]domain.Bridger),
	}

	for _, useCase := range useCases {
		switch useCase.(type) {
		case domain.ArticleUseCase:
			blog.useCases[config.ARTICLE_COLLECTION] = useCase
		}
	}

	blogPB.RegisterBlogServiceServer(server, blog)
}
