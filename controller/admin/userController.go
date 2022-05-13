package admin

import (
	"gin/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) Index(c *gin.Context) {
	services.UserService{}.Index(c)
}
func (con UserController) Add(c *gin.Context) {
	services.UserService{}.Add(c)
}
func (con UserController) Update(c *gin.Context) {
	services.UserService{}.Update(c)
}
func (con UserController) Delete(c *gin.Context) {
	services.UserService{}.Delete(c)
}
