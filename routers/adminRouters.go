package routers

import (
	"gin/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {

	adminRouters := r.Group("/admin")
	{
		// admin.UserIndex() 表示执行方法，admin.UserIndex 表示注册方法
		adminRouters.GET("/", admin.IndexController{}.Index)

		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.POST("/user/add", admin.UserController{}.Add)
		adminRouters.PUT("/user/update", admin.UserController{}.Update)
		adminRouters.DELETE("/user/delete", admin.UserController{}.Delete)

		adminRouters.GET("/user1", admin.User1Controller{}.Index)
		adminRouters.POST("/user1/add", admin.User1Controller{}.Add)
		adminRouters.PUT("/user1/update", admin.User1Controller{}.Update)
		adminRouters.DELETE("/user1/delete", admin.User1Controller{}.Delete)
		// :id 相当于一个变量，函数中用c.param()接收解析
		//adminRouters.DELETE("/user1/delete/:id", admin.User1Controller{}.Delete)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/update", admin.ArticleController{}.Update)
		adminRouters.GET("/article/delete", admin.ArticleController{}.Delete)
	}
}
