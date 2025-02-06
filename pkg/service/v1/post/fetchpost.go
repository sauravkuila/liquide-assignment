package post

import (
	"liquide-assignment/pkg/config"
	"liquide-assignment/pkg/dto"
	e "liquide-assignment/pkg/errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (obj *postService) GetPost(c *gin.Context) {
	var (
		request  dto.GetPostRequest
		response dto.GetPostResponse
	)
	if err := c.BindUri(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to get post info"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//get the post info from db
	postInfo, err := obj.dbObj.GetPost(c, request.PostId)
	if err != nil {
		log.Printf("unable to get post info. Error:%s", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Status = true
			response.Errors = append(response.Errors, *e.ErrorInfo[e.NoDataFound])
			response.Message = "post not found"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Errors = append(response.Errors, *e.ErrorInfo[e.GetDBError])
		response.Message = "failed to get post info"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	postInf := postInfo.ToPostInfo()
	response.Data = &postInf
	response.Status = true
	response.Message = "post info fetched successfully"
	c.JSON(http.StatusOK, response)
}

func (obj *postService) GetAllPostsForUser(c *gin.Context) {
	var (
		request  dto.GetAllUserPostRequest
		response dto.GetAllUserPostResponse
	)
	if err := c.BindQuery(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to get post info"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//get the post info from db
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PageSize < 1 {
		request.PageSize = 10
	}
	offset := (request.Page - 1) * request.PageSize
	limit := request.PageSize
	posts, totalRecords, err := obj.dbObj.GetUserPosts(c, c.GetInt64(config.USERID), limit, offset)
	if err != nil {
		log.Printf("unable to get post info. Error:%s", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Status = true
			response.Errors = append(response.Errors, *e.ErrorInfo[e.NoDataFound])
			response.Message = "post not found"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Errors = append(response.Errors, *e.ErrorInfo[e.GetDBError])
		response.Message = "failed to get post info"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Data = posts
	response.Page = request.Page
	response.PageSize = request.PageSize
	response.TotalCount = totalRecords
	response.Status = true
	response.Message = "post info fetched successfully"
	c.JSON(http.StatusOK, response)
}
