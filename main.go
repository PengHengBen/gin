package main

import (
	"fmt"
	"gin/routers"
	"github.com/gin-gonic/gin"
	"time"
)

// 路由中间件
func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-middleware function")
	// 调用请求剩余处理程序
	c.Next()
	// 中止调用请求处理程序
	//c.Abort()
	fmt.Println("2-middleware function")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func main() {
	// 创建默认路由引擎
	r := gin.Default()
	// 自定义模板函数 注意要把这个函数放在加载模板前
	//r.SetFuncMap(template.FuncMap{
	//	"UnixToTime": models.UnixToTime(time.Now().Unix()),
	//})
	// 配置模板文件路径
	r.LoadHTMLGlob("templates/**/*")
	// 全局中间件
	//r.Use(initMiddleware)
	// 加载路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	// 配置路由中间件
	//r.GET("/middleware", func(c *gin.Context) {
	//	fmt.Println("this is middleware home page")
	//	//time.Sleep(time.Second)
	//	c.String(http.StatusOK, "middleware home page")
	//})
	r.Run()
}
