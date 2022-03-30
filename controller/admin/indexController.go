package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	// middleware文件夹下init.go文件中设置的username
	//username, _ := c.Get("username")
	//fmt.Println(username)
	//
	//// 空接口类型转换为string类型
	//v, ok := username.(string)
	//// 类型断言
	//if ok {
	//	c.String(http.StatusOK, "用户列表--"+v)
	//} else {
	//	c.String(http.StatusOK, "获取用户失败")
	//}
	c.String(http.StatusOK, "admin index list controller struct")
}
