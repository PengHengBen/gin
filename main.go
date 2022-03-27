package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// timestamp convert to date
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()
	// 配置模板文件路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		// notice: this needs LoadHTMLGlob()
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "this is goods backend data",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{ // == gin.H
			"success": true,
			"msg":     "hello gin",
		})
	})
	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "post data")
	})
	r.PUT("/update", func(c *gin.Context) {
		c.String(http.StatusOK, "update data")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete data")
	})

	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{})
	})
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	r.Run()
}
