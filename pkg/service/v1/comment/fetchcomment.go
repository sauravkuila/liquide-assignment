package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *commentService) GetComment(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *commentService) GetAllComment(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
