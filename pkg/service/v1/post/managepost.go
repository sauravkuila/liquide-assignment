package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *postService) CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *postService) EditPost(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *postService) DeletePost(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
