package home

import (
	"fmt"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println(utils.GetDate())
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "Hello gin, this is home page",
		"t":   utils.GetDate(), // 时间戳的单位是秒
	})
}
func (con DefaultController) News(c *gin.Context) {
	c.String(http.StatusOK, "default controller news struct")
}
