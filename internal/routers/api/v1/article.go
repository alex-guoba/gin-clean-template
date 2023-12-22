package v1

import (
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"

	"github.com/alex-guoba/gin-clean-template/internal/service"
	"github.com/alex-guoba/gin-clean-template/pkg/app"
	"github.com/alex-guoba/gin-clean-template/pkg/convert"
	"github.com/alex-guoba/gin-clean-template/pkg/errcode"
	"github.com/alex-guoba/gin-clean-template/pkg/logger"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (Article) checkParams(c *gin.Context, param any, response *app.Response) error {
	if err := c.ShouldBind(param); err != nil {
		logger.WithTrace(c).Errorf("params errs: %v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return err
	}
	return nil
}

// @Summary 创建文章
// @Produce json
// @Param tag_id body string true "标签ID"
// @Param title body string true "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string true "封面图片地址"
// @Param content body string true "文章内容"
// @Param created_by body int true "创建者"
// @Param state body int false "状态"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [post].
func (art Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	if art.checkParams(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 获取单个文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [get].
func (art Article) Get(c *gin.Context) {
	param := service.ArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	if art.checkParams(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
}

// @Summary 获取多个文章
// @Produce json
// @Param name query string false "文章名称"
// @Param tag_id query int false "标签ID"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles [get].
func (art Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	if art.checkParams(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles, totalRows, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	for _, article := range articles {
		log.Info(article)
	}
	log.Info("total num: ", totalRows)

	response.ToResponseList(articles, totalRows)
}

// @Summary 更新文章
// @Produce json
// @Param tag_id body string false "标签ID"
// @Param title body string false "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string false "封面图片地址"
// @Param content body string false "文章内容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put].
func (art Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	if art.checkParams(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary 删除文章
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete].
func (art Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	if art.checkParams(c, &param, response) != nil {
		return
	}

	svc := service.NewArticleService(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		logger.WithTrace(c).Errorf("svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}

	response.ToResponse(gin.H{})
}
