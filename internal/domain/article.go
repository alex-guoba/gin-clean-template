package domain

import (
	"context"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/entity"

	"gorm.io/gorm"
)

type ArticleDomain struct {
	ctx          context.Context
	articleDao   dao.ArticleDao
	tagDao       dao.TagDao
	artileTagDao dao.ArticleTagDao
}

func NewArticleDomain(ctx context.Context, db *gorm.DB) *ArticleDomain {
	d := &ArticleDomain{ctx: ctx}
	d.tagDao = dao.NewTagDao(db)
	d.articleDao = dao.NewArticleDaoDB(db)
	d.artileTagDao = dao.NewArticleTagDao(db)
	return d
}

func (d *ArticleDomain) GetArticle(id uint32, state uint8) (*entity.ArticleEntity, error) {
	// Query article info
	article, err := d.articleDao.GetArticle(id, state)
	if err != nil {
		return nil, err
	}

	// Query state id
	articleTag, err := d.artileTagDao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	// Query tag info
	tag, err := d.tagDao.GetTag(articleTag.TagID, dao.StateOpen)
	if err != nil {
		return nil, err
	}

	// TODO: convert to entry object
	return &entity.ArticleEntity{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageURL: article.CoverImageURL,
		State:         article.State,
		Tag: &entity.TagEntity{
			ID:         tag.ID,
			Name:       tag.Name,
			State:      tag.State,
			CreatedBy:  tag.CreatedBy,
			ModifiedBy: tag.ModifiedBy,
		},
	}, nil
}

func (d *ArticleDomain) countArticleByTagID(id uint32) (int, error) {
	cnt, err := d.articleDao.CountArticleListByTagID(id, 1)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (d *ArticleDomain) GetArticleList(id uint32, state uint8, page, pageSize int) ([]*entity.ArticleEntity, int, error) {
	// Query article count
	cnt, err := d.articleDao.CountArticleListByTagID(id, state)
	if err != nil {
		return nil, 0, err
	}

	// Query article list
	artileTags, err := d.articleDao.GetArticleListByTagID(id, state, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*entity.ArticleEntity
	for _, row := range artileTags {
		articleList = append(articleList, &entity.ArticleEntity{
			ID:            row.Article.ID,
			Title:         row.Article.Title,
			Desc:          row.Article.Desc,
			Content:       row.Article.Content,
			CoverImageURL: row.Article.CoverImageURL,
			Tag: &entity.TagEntity{
				ID:   row.Tag.ID,
				Name: row.Tag.Name,
			},
		})
	}

	return articleList, int(cnt), nil
}

func (d *ArticleDomain) CreateArticle(title string, desc string, content string, image string,
	state uint8, createdBy string, tagID uint32) error {
	// Insert article
	article, err := d.articleDao.CreateArticle(title, desc, content, image, state, createdBy)
	if err != nil {
		return err
	}

	// Insert article tag relation
	return d.artileTagDao.CreateArticleTag(article.ID, tagID, createdBy)
}

func (d *ArticleDomain) UpdateArticle(artitleID uint32, title string, desc string, content string, image string,
	state uint8, modifiedBy string, tagID uint32) error {
	// Update article
	if err := d.articleDao.UpdateArticle(artitleID, title, desc, content,
		image, state, modifiedBy); err != nil {
		return err
	}

	// Update article tag relation
	return d.artileTagDao.UpdateArticleTag(artitleID, tagID, modifiedBy)
}

func (d *ArticleDomain) DeleteArticle(id uint32) error {
	// Delete article
	if err := d.articleDao.DeleteArticle(id); err != nil {
		return err
	}

	// Delete article tag relation
	return d.artileTagDao.DeleteArticleTag(id)
}
