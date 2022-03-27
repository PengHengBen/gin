package controller

import "pkg/mod/github.com/gin-gonic/gin@v1.7.7"

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "user list")
}

func (con UserController) add(c *gin.Context) {
	c.String(200, "add user")
}

func (con UserController) Update(c *gin.Context) {
	c.String(200, "Update user")
}

func (con UserController) Delete(c *gin.Context) {
	c.String(200, "delete user")
}
