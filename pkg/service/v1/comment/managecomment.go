package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *commentService) CreateComment(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *commentService) EditComment(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *commentService) DeleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
