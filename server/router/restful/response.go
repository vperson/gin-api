package restful

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	c *gin.Context
}

func New(c *gin.Context) *Response {
	return &Response{c: c}
}

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *Response) Ok(data interface{}) {
	resp := ResponseData{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: data,
	}

	r.c.JSON(http.StatusOK, resp)
}

func (r *Response) StatusBadRequest(errMsg string) {
	msg := "parameter error"
	if errMsg != "" {
		msg = errMsg
	}

	resp := ResponseData{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: nil,
	}

	r.c.JSON(http.StatusBadRequest, resp)
}

func (r *Response) StatusBadRequestWithCode(errMsg string, code int) {
	msg := "parameter error"
	respCode := http.StatusInternalServerError
	if errMsg != "" {
		msg = errMsg
	}
	if code != 0 {
		respCode = code
	}

	resp := ResponseData{
		Code: respCode,
		Msg:  msg,
		Data: nil,
	}

	r.c.JSON(http.StatusBadRequest, resp)
}

func (r *Response) StatusInternalServerError(errMsg string) {
	msg := "failed"
	if errMsg != "" {
		msg = errMsg
	}

	resp := ResponseData{
		Code: http.StatusInternalServerError,
		Msg:  msg,
		Data: nil,
	}

	r.c.JSON(http.StatusInternalServerError, resp)
}
