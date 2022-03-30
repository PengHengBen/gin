package home

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	fmt.Println(models.GetDate())
	c.HTML(http.StatusOK, "default/index.html", gin.H{
		"msg": "Hello gin, this is home page",
		"t":   models.GetDate(), // 时间戳的单位是秒
	})
}
func (con DefaultController) News(c *gin.Context) {
	c.String(http.StatusOK, "default controller news struct")
}
