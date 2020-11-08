package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 0
	FAILED  int = 1
)

// 普通成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})
}

// 普通失败返回
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"smg":  v,
	})
}
