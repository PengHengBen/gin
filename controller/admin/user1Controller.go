package admin

import (
	"gin/services"
	"github.com/gin-gonic/gin"
)

type User1Controller struct{}

func (con User1Controller) Index(c *gin.Context) {
	services.User1Service{}.Index(c)
}
func (con User1Controller) Add(c *gin.Context) {
	services.User1Service{}.Add(c)
}
func (con User1Controller) Update(c *gin.Context) {
	services.User1Service{}.Update(c)
}
func (con User1Controller) Delete(c *gin.Context) {
	services.User1Service{}.Delete(c)
}
