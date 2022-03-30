package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	// 分组中间件在路由模板配置
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	// 设置键值，在其他模块使用c.Get获取，这次使用在admin indexController中
	c.Set("username", "zhangsan")

	// 定义一个goroutine
	// gin中间件使用goroutine：当在中间件或handler中启动新的goroutine时，
	// 不能使用原始上下文(c *gin.Context)，必须使用其只读副本(c.Copy())
	// 程序响应时间不受影响，只是打印会内容会延迟两秒输出
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
