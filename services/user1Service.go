package services

import (
	"fmt"
	"gin/initialize"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User1Service struct {
	BaseService
}

type User1Bind struct {
	Id     int    `form:"id"`
	Name   string `form:"name"`
	Gender string `form:"gender"`
	Age    int    `form:"age"`
}

func bind(c *gin.Context) (user1 models.User1, err error) {
	p := User1Bind{}
	if err = c.BindJSON(&p); err != nil {
		// 使用这个会中断退出函数，后边的语句不会执行
		//log.Fatal(err)
		fmt.Println(err)
		return user1, err
	}
	user1 = models.User1{
		Id:             p.Id,
		Name:           p.Name,
		Gender:         p.Gender,
		Age:            p.Age,
		SubmissionDate: time.Now(),
	}
	return user1, err
}

func findOne(id int) bool {
	user1 := models.User1{}
	if err := initialize.DB.First(&user1, id).Error; err != nil {
		return false
	}
	return true
}

func (con User1Service) Index(c *gin.Context) {
	// 查询数据库
	user1List := []models.User1{}
	initialize.DB.Find(&user1List)
	c.JSON(http.StatusOK, gin.H{
		"result": user1List,
	})

	// GET method 绑定结构体传值（url后边加参数）
	//p := User1Bind{}
	//if err:=c.ShouldBindQuery(&p); err !=nil{
	//	log.Fatal(err)
	//  return
	//}
	//fmt.Println(p.Id,p.Name,p.Gender,p.Age,p)
	//con.success(c)
}
func (con User1Service) Add(c *gin.Context) {
	//user := c.PostForm("user")
	//gender := c.PostForm("gender")
	//age := c.PostForm("age")
	//
	//user1 := models.User1{
	//	Name:           user,
	//	Gender:         gender,
	//	Age:            age,
	//	SubmissionDate: utils.GetDate(),
	//}
	//initialize.DB.Create(&user1)

	//p := User1Bind{}
	// POST method 请求体的form-data数据并绑定到结构体
	//if err:=c.ShouldBind(&p); err !=nil{
	//	log.Fatal(err)
	//    return
	//}
	// POST method 请求体的json数据并绑定到结构体
	//if err := c.BindJSON(&p); err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//fmt.Println(p.Id, p.Name, p.Gender, p.Age, p)
	//user1 := models.User1{
	//	Name:           p.Name,
	//	Gender:         p.Gender,
	//	Age:            p.Age,
	//	SubmissionDate: time.Now(),
	//}

	user1, err := bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -2,
			"message": "error params",
		})
		return
	}
	tx := initialize.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	if tx.Error != nil {
		fmt.Println("事物开启异常")
	}
	if err = tx.Create(&user1).Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	// 等于c.JSON(http.StatusOK, map[string]interface{}{}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "ok",
		"data":    user1.Name + " add success",
	})
}
func (con User1Service) Update(c *gin.Context) {
	user1, err := bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -2,
			"message": "error params",
		})
		return
	}

	if !findOne(user1.Id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -3,
			"message": "id dose not exist",
		})
		return
	}

	tx := initialize.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	if tx.Error != nil {
		fmt.Println("事物开启异常")
	}
	// 使用结构体更新数据时，不传值就不更新，也就是只更新传递的值，其余值保持默认，但是使用结构体时，定义的字段key是数据库模板里的
	// 判断err时，需要在更新的语句后边添加.Error 直接返回的是事务类型，.Error返回的是error
	err = tx.Model(&user1).Where("id = ?", user1.Id).Updates(models.User1{Name: user1.Name, Gender: user1.Gender, Age: user1.Age, SubmissionDate: user1.SubmissionDate}).Error
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		c.JSON(500, gin.H{
			"code":    -1,
			"message": "update failed",
		})
	}
	tx.Commit()
	// 等于c.JSON(http.StatusOK, map[string]interface{}{}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "update success",
		"data":    user1.Id,
	})

	//initialize.DB.Model(&user1).Where("id = ?", id).Updates(map[string]interface{}{"name": user, "gender": gender, "age": age, "submission_date": utils.GetDate()})
	//c.String(http.StatusOK, "update user1 success")
}
func (con User1Service) Delete(c *gin.Context) {
	//id := c.PostForm("id")
	//user1 := models.User1{}

	//initialize.DB.Where("id = ?", id).Delete(&user1)
	// 接收url中传递的变量id
	//params := c.Param("id")
	//fmt.Println(reflect.TypeOf(params))
	//fmt.Println(params)

	user1, err := bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -2,
			"message": "error params",
		})
		return
	}

	if !findOne(user1.Id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -3,
			"message": "id dose not exist",
		})
		return
	}

	tx := initialize.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	err = tx.Where("id = ?", user1.Id).Delete(&user1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "delete success",
		"data":    user1.Id,
	})
}
