package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/alex-guoba/gin-clean-template/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

type ListResponse struct {
	// 列表数据
	List  any   `json:"list"`
	Pager Pager `json:"pager"`
}

type MapResponse map[string]string

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data any) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// func (r *Response) ToResponseList(list any, totalRows int, page int, pageSize int) {
// 	r.Ctx.JSON(http.StatusOK, gin.H{
// 		"list": list,
// 		"pager": Pager{
// 			Page:      page,
// 			PageSize:  pageSize,
// 			TotalRows: totalRows,
// 		},
// 	})
// }

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code, "msg": err.Msg}
	details := err.Details
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
