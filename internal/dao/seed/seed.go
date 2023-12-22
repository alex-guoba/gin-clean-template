package seed

import (
	"github.com/alex-guoba/gin-clean-template/internal/dao"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

// Seed database with test data.

func Seed(engine *gorm.DB, count int) error {
	return seedArticleTag(engine, count)
}

func seedArticleTag(engine *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		model := dao.Model{CreatedBy: "gofakeit"}

		article := dao.ArticleModel{
			Title:         gofakeit.BookTitle(),
			Desc:          gofakeit.ProductDescription(),
			Content:       gofakeit.JobDescriptor(),
			CoverImageURL: gofakeit.URL(),
			State:         1,
			Model:         &model,
		}
		if err := engine.Create(&article).Error; err != nil {
			return err
		}

		tag := dao.TagModel{
			Name:  gofakeit.Gamertag(),
			State: 1,
			Model: &model,
		}
		if err := engine.Create(&tag).Error; err != nil {
			return err
		}

		at := dao.ArticleTagModel{
			Model:     &model,
			ArticleID: article.ID,
			TagID:     tag.ID,
		}
		if err := engine.Create(&at).Error; err != nil {
			return err
		}
	}

	return nil
}
