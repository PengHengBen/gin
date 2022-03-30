package admin

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User1Controller struct {
	BaseController
}

// 开启事物
//tx := global.DBLink.Begin()

func (con User1Controller) Index(c *gin.Context) {
	//c.String(http.StatusOK,"admin user list controller struct")
	//con.success(c)

	// 查询数据库
	user1List := []models.User1{}
	models.DB.Find(&user1List)
	c.JSON(http.StatusOK, gin.H{
		"result": user1List,
	})
}
func (con User1Controller) Add(c *gin.Context) {
	user1 := models.User1{
		Name:           "xiaosun",
		Gender:         "male",
		Age:            22,
		SubmissionDate: time.Now(),
	}
	models.DB.Create(&user1)
	fmt.Println(user1)
	c.String(http.StatusOK, "add user1 success")
}
func (con User1Controller) Update(c *gin.Context) {
	user1 := models.User1{}
	// 更新单列
	models.DB.Model(&user1).Where("id = ?", 2).Update("submission_date", time.Now())
	c.String(http.StatusOK, "update user1 success")
}
func (con User1Controller) Delete(c *gin.Context) {
	user1 := models.User1{Id: 4}
	models.DB.Delete(&user1)
	c.String(http.StatusOK, "delete data success")
}
