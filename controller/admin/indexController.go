package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	c.String(http.StatusOK, "admin index list controller struct")
}
