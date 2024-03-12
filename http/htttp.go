package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Responses(c *gin.Context, code int64, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}
