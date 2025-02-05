package vote

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *voteService) UpVote(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}

func (obj *voteService) DownVote(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
