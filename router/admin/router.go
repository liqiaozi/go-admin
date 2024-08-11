package admin

import "github.com/gin-gonic/gin"

var (
	noCheckRouter = make([]func(*gin.RouterGroup), 0)
	checkRouter   = make([]func(*gin.RouterGroup), 0)
)

func InitRouter(r *gin.Engine) {
	// 处理无需认证的路由
	handleNoCheckRouter(r)
	// 处理需要认证的路由
	handleCheckRouter(r)
}

func handleCheckRouter(r *gin.Engine) {
	group := r.Group("/api/v1/admin")
	for _, f := range noCheckRouter {
		f(group)
	}
}

func handleNoCheckRouter(r *gin.Engine) {
	group := r.Group("/api/v1/admin")
	for _, f := range checkRouter {
		f(group)
	}
}
