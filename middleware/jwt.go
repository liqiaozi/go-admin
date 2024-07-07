package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/component"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.NoAuthWithMsg(c, "未登录或非法访问")
			c.Abort()
			return
		}
		// 是否在黑名单库中

		claims, err := component.JwtService.ParseToken(token)
		if err != nil {
			if errors.Is(err, component.TokenExpired) {
				response.NoAuthWithMsg(c, "授权已过期")
			} else {
				response.NoAuthWithMsg(c, err.Error())
			}
			c.Abort()
			return
		}

		// Token 发布者校验
		if claims.Issuer != component.Issuer {
			response.NoAuth(c)
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("id", claims.Id)
	}

}
