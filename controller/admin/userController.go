package admin

import (
	"fmt"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	//c.String(http.StatusOK,"admin user list controller struct")
	//con.success(c)

	// 查询数据库
	userList := []models.User{}
	models.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}
func (con UserController) Add(c *gin.Context) {
	user := models.User{
		Name:           "xiaoli",
		Gender:         "male",
		Age:            22,
		SubmissionDate: time.Now(),
	}
	models.DB.Create(&user)
	fmt.Println(user)
	c.String(http.StatusOK, "add user success")
}
func (con UserController) Update(c *gin.Context) {
	// 查询id等于1的数据
	user := models.User{Id: 1}
	// 更新数据 要写上所有字段，不然字段会被更新为空值
	user.Name = "xiaowang"
	user.Gender = "female"
	user.Age = 33
	user.SubmissionDate = time.Now()
	models.DB.Save(&user)
	c.String(http.StatusOK, "update user success")
}
func (con UserController) Delete(c *gin.Context) {
	//user := models.User{Id: 2}
	//models.DB.Delete(&user)
	user := models.User{}
	models.DB.Where("username=?", "xiaowang").Delete(&user)
	c.String(http.StatusOK, "delete data success")
}
