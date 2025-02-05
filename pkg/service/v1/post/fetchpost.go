package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *postService) GetPost(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *postService) GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
