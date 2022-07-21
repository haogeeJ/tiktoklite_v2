package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回错误信息
func ErrResponse(c *gin.Context, statusCode int32, statusMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": statusCode,
		"status_msg":  statusMsg,
	})
}
