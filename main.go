package main

import (
	"gin/initialize"
	"gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置文件
	initialize.Init()
	// 初始化数据库
	initialize.InitDb()
	// 初始化表
	initialize.InitTable()
	// 创建默认路由引擎
	r := gin.Default()
	// 配置模板文件路径
	r.LoadHTMLGlob("templates/**/*")
	// 加载路由
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)
	r.Run(":" + initialize.V.Get("app.port").(string))
}
