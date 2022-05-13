package services

import (
	"gin/initialize"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService struct {
	BaseService
}

func (con UserService) Index(c *gin.Context) {
	// 查询数据库
	userList := []models.User{}
	initialize.DB.Find(&userList)
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
	//con.success(c)
}
func (con UserService) Add(c *gin.Context) {
	//user := models.User{
	//	Name:           "xiaosun",
	//	Gender:         "male",
	//	Age:            "22",
	//	SubmissionDate: utils.GetDate(),
	//}
	//initialize.DB.Create(&user)
	//fmt.Println(user)
	c.String(http.StatusOK, "add user1 success")
}
func (con UserService) Update(c *gin.Context) {
	//user := models.User{}
	//// 更新单列
	//initialize.DB.Model(&user).Where("id = ?", 2).Update("submission_date", time.Now())
	c.String(http.StatusOK, "update user1 success")
}
func (con UserService) Delete(c *gin.Context) {
	//user := models.User{Id: 2}
	//initialize.DB.Delete(&user)
	c.String(http.StatusOK, "delete data success")
}
