package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.String(http.StatusOK, "api controller index struct")
}
func (con ApiController) Userlist(c *gin.Context) {
	c.String(http.StatusOK, "api controller user list struct")
}
func (con ApiController) Plist(c *gin.Context) {
	c.String(http.StatusOK, "api controller plist struct")
}
