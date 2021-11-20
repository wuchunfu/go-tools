package responsex

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	// 状态码
	Code int `json:"code" example:"200"`
	// 消息
	Msg string `json:"msg"`
	// 数据集
	Data interface{} `json:"data"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

type PageResponse struct {
	// 状态码
	Code int `json:"code" example:"200"`
	// 消息
	Msg string `json:"msg"`
	// 数据集
	Data Page `json:"data"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}

// Error 失败数据处理
func Error(ctx *gin.Context, code int, err error, msg string) {
	var res Response
	res.Msg = err.Error()
	if msg != "" {
		res.Msg = msg
	}
	ctx.JSON(http.StatusOK, res.ReturnError(code))
}

// OK 通常成功数据处理
func OK(ctx *gin.Context, data interface{}, msg string) {
	var res Response
	res.Data = data
	if msg != "" {
		res.Msg = msg
	}
	ctx.JSON(http.StatusOK, res.ReturnOK())
}

// PageOK 分页数据处理
func PageOK(ctx *gin.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res PageResponse
	res.Data.List = result
	res.Data.Count = count
	res.Data.PageIndex = pageIndex
	res.Data.PageSize = pageSize
	if msg != "" {
		res.Msg = msg
	}
	ctx.JSON(http.StatusOK, res.ReturnOK())
}

// Custom 兼容函数
func Custom(ctx *gin.Context, data gin.H) {
	ctx.JSON(http.StatusOK, data)
}
