package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseService struct{}

func (con BaseService) success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
func (con BaseService) error(c *gin.Context) {
	c.String(http.StatusBadRequest, "error")
}
