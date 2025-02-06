package comment

import (
	"liquide-assignment/pkg/dto"
	e "liquide-assignment/pkg/errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (obj *commentService) GetComment(c *gin.Context) {
	var (
		request  dto.GetCommentRequest
		response dto.GetCommentResponse
	)
	if err := c.BindUri(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to fetch comment"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//fetch the comment from db

	c.JSON(http.StatusOK, "api in development")
}

func (obj *commentService) GetAllComment(c *gin.Context) {
	// var (
	// 	request  dto.GetAllCommentRequest
	// 	response dto.GetAllCommentResponse
	// )
	c.JSON(http.StatusOK, "api in development")
}

func (obj *commentService) GetAllReplies(c *gin.Context) {
	c.JSON(http.StatusOK, "api in development")
}
