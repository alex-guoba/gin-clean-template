package service

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/domain"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
	"github.com/alex-guoba/gin-clean-template/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// demo requst: curl -v -X POST "http://localhost:8000/api/v1/articles?tag_id=2&title=test&desc=testdesc&content=testcontent&cover_image_url=https://www.google.com&created_by=test&state=1"
type CreateArticleRequest struct {
	TagID         uint32 `form:"tag_id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=2,max=100"`
	Desc          string `form:"desc" binding:"required,min=2,max=255"`
	Content       string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,url"`
	CreatedBy     string `form:"created_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	TagID         uint32 `form:"tag_id" binding:"omitempty,gte=1"`
	Title         string `form:"title" binding:"omitempty,min=2,max=100"`
	Desc          string `form:"desc" binding:"omitempty,min=2,max=255"`
	Content       string `form:"content" binding:"omitempty,min=2,max=4294967295"`
	CoverImageUrl string `form:"cover_image_url" binding:"omitempty,url"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
	State         uint8  `form:"state,default=1" binding:"omitempty,oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// service
type ArticleService struct {
	ctx    context.Context
	domain *domain.ArticleDomain
}

func NewArticleService(ctx context.Context) *ArticleService {
	return &ArticleService{ctx: ctx, domain: domain.NewArticleDomain(ctx)}
}

func (svc *ArticleService) GetArticle(param *ArticleRequest) (*entity.ArticleEntity, error) {
	article, err := svc.domain.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (svc *ArticleService) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*entity.ArticleEntity, int, error) {
	articleList, cnt, err := svc.domain.GetArticleList(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	return articleList, cnt, nil
}

func (svc *ArticleService) CreateArticle(param *CreateArticleRequest) error {
	err := svc.domain.CreateArticle(param.Title, param.Desc,
		param.Content, param.CoverImageUrl, param.State, param.CreatedBy, param.TagID)
	if err != nil {
		return err
	}

	return nil
}

func (svc *ArticleService) UpdateArticle(param *UpdateArticleRequest) error {
	if err := svc.domain.UpdateArticle(param.ID, param.Title, param.Desc, param.Content,
		param.CoverImageUrl, param.State, param.ModifiedBy, param.TagID); err != nil {
		return err
	}

	return nil
}

func (svc *ArticleService) DeleteArticle(param *DeleteArticleRequest) error {
	if err := svc.domain.DeleteArticle(param.ID); err != nil {
		return err
	}

	return nil
}