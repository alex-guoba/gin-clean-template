package domain

import (
	"context"
	"reflect"
	"testing"

	"github.com/alex-guoba/gin-clean-template/internal/dao"
	"github.com/alex-guoba/gin-clean-template/internal/dao/mocks"
	"github.com/alex-guoba/gin-clean-template/internal/entity"
)

// func newArticleDomainMock() *ArticleDomain {
// 	d := &ArticleDomain{ctx: context.Background()}
// 	d.tagDao = &mocks.TagDao{}
// 	d.articleDao = &mocks.ArticleDao{}
// 	d.artileTagDao = &mocks.ArticleTagDao{}
// 	return d
// }

func newModelsMock() (*dao.ArticleModel, *dao.TagModel, *dao.ArticleTagModel) {
	article := &dao.ArticleModel{
		Title:         "title",
		Desc:          "desc",
		Content:       "content",
		CoverImageURL: "image",
		State:         1,
		Model: &dao.Model{
			ID:         1,
			CreatedBy:  "author",
			ModifiedBy: "author",
			IsDel:      0,
		},
	}
	tag := &dao.TagModel{
		Name:  "tag",
		State: 1,
		Model: &dao.Model{
			ID:         2,
			CreatedBy:  "author",
			ModifiedBy: "author",
			IsDel:      0,
		},
	}
	articleTag := &dao.ArticleTagModel{
		ArticleID: 1,
		TagID:     2,
		Model: &dao.Model{
			ID:         1,
			CreatedBy:  "author",
			ModifiedBy: "author",
			IsDel:      0,
		},
	}
	return article, tag, articleTag
}

func TestArticleDomain_GetArticle(t *testing.T) {
	// mock data
	article, tag, articleTag := newModelsMock()

	articleDao := &mocks.ArticleDao{}
	articleDao.On("GetArticle", article.ID, article.State).Return(*article, nil)

	tagDao := &mocks.TagDao{}
	tagDao.On("GetTag", tag.ID, uint8(dao.StateOpen)).Return(*tag, nil)

	artileTagDao := &mocks.ArticleTagDao{}
	artileTagDao.On("GetArticleTagByAID", article.ID).Return(*articleTag, nil)

	d := &ArticleDomain{
		ctx:          context.Background(),
		articleDao:   articleDao,
		tagDao:       tagDao,
		artileTagDao: artileTagDao,
	}

	got, err := d.GetArticle(article.ID, article.State)
	if err != nil {
		t.Errorf("ArticleDomain.GetArticle() error = %v", err)
		return
	}
	t.Logf("%v", got)
}

func TestArticleDomain_countArticleByTagID(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			got, err := d.countArticleByTagID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.countArticleByTagID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ArticleDomain.countArticleByTagID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArticleDomain_GetArticleList(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		id       uint32
		state    uint8
		page     int
		pageSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.ArticleEntity
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			got, got1, err := d.GetArticleList(tt.args.id, tt.args.state, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.GetArticleList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleDomain.GetArticleList() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ArticleDomain.GetArticleList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArticleDomain_CreateArticle(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		title     string
		desc      string
		content   string
		image     string
		state     uint8
		createdBy string
		tagID     uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			if err := d.CreateArticle(tt.args.title, tt.args.desc, tt.args.content, tt.args.image, tt.args.state, tt.args.createdBy, tt.args.tagID); (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleDomain_UpdateArticle(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		artitleID  uint32
		title      string
		desc       string
		content    string
		image      string
		state      uint8
		modifiedBy string
		tagID      uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			if err := d.UpdateArticle(tt.args.artitleID, tt.args.title, tt.args.desc, tt.args.content, tt.args.image, tt.args.state, tt.args.modifiedBy, tt.args.tagID); (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArticleDomain_DeleteArticle(t *testing.T) {
	type fields struct {
		ctx          context.Context
		articleDao   dao.ArticleDao
		tagDao       dao.TagDao
		artileTagDao dao.ArticleTagDao
	}
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ArticleDomain{
				ctx:          tt.fields.ctx,
				articleDao:   tt.fields.articleDao,
				tagDao:       tt.fields.tagDao,
				artileTagDao: tt.fields.artileTagDao,
			}
			if err := d.DeleteArticle(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ArticleDomain.DeleteArticle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
