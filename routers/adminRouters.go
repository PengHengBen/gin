package routers

import (
	"gin/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	// 配置分组中间件
	//adminRouters:=r.Group("/admin",middleware.InitMiddleware)
	adminRouters := r.Group("/admin")
	{
		// admin.UserIndex() 表示执行方法，admin.UserIndex 表示注册方法
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/update", admin.UserController{}.Update)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)

		adminRouters.GET("/user1", admin.User1Controller{}.Index)
		adminRouters.GET("/user1/add", admin.User1Controller{}.Add)
		adminRouters.GET("/user1/update", admin.User1Controller{}.Update)
		adminRouters.GET("/user1/delete", admin.User1Controller{}.Delete)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/update", admin.ArticleController{}.Update)
		adminRouters.GET("/article/delete", admin.ArticleController{}.Delete)
	}
}
