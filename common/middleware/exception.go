package middleware

import (
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"net/http"
	"runtime/debug"
)

// CustomExceptionHandler 自定义全局异常处理中间件
func CustomExceptionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 捕获 panic
				logger.Log.Errorf("Recovered from panic: %v, detail: \n %v", err, string(debug.Stack()))
				// 判断是否是自定义错误类型
				if customErr, ok := err.(errors.CustomError); ok {
					c.JSON(http.StatusInternalServerError, response.FailByError(customErr))
				} else {
					c.JSON(http.StatusInternalServerError, response.Fail())
				}
				// 终止请求处理链
				c.Abort()
			}
		}()
		c.Next()
	}
}
