package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct{}

func (con ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "admin article list controller struct")
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "admin add article controller struct")
}
func (con ArticleController) Update(c *gin.Context) {
	c.String(http.StatusOK, "admin update article controller struct")
}
func (con ArticleController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "admin delete article controller struct")
}
