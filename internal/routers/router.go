package routers

import (
	v1 "github.com/alex-guoba/gin-clean-template/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		// Create tag
		apiv1.POST("/tags", tag.Create)
		// Delete tag
		apiv1.DELETE("/tags/:id", tag.Delete)
		// Update tab
		apiv1.PUT("/tags/:id", tag.Update)
		// Get tag
		apiv1.GET("/tags", tag.List)

		// Create artile with tags
		apiv1.POST("/articles", article.Create)
		// Get aritle detail by id
		apiv1.GET("/articles/:id", article.Get)
		// Get article list and total count
		apiv1.GET("/articles", article.List)
		// Update article detail by id
		apiv1.PUT("/articles/:id", article.Update)
		// Delete article by id
		apiv1.DELETE("/articles/:id", article.Delete)

		// Add other router if necessary
	}

	return r
}
