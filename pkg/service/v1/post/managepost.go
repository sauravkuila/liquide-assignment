package post

import (
	"database/sql"
	"liquide-assignment/pkg/blog"
	"liquide-assignment/pkg/config"
	"liquide-assignment/pkg/dto"
	"log"
	"net/http"

	e "liquide-assignment/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (obj *postService) CreatePost(c *gin.Context) {
	var (
		request  dto.CreatePostRequest
		response dto.CreatePostResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to create post"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//add the entry into db
	dbPost := dto.DbPost{
		UserId:  sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		Content: sql.NullString{String: request.Content, Valid: true},
	}
	postId, err := obj.dbObj.CreatePost(c, dbPost)
	if err != nil {
		log.Printf("failed to create post. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to create post"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	//update the blog about post creation
	go blog.NewBlogObject(obj.redisObj).AddPost(c, dto.Post{
		PostId:  postId,
		UserId:  c.GetInt64(config.USERID),
		Content: request.Content,
	})

	response.Status = true
	response.Data = &dto.Post{
		PostId:   postId,
		Content:  request.Content,
		UserId:   c.GetInt64(config.USERID),
		UserName: c.GetString(config.USERNAME),
	}
	response.Message = "successfully created post"

	c.JSON(http.StatusOK, response)
}

func (obj *postService) EditPost(c *gin.Context) {
	var (
		request  dto.EditPostRequest
		response dto.EditPostResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to edit post"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//add the entry into db
	dbPost := dto.DbPost{
		PostId:  sql.NullInt64{Int64: request.PostId, Valid: true},
		UserId:  sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		Content: sql.NullString{String: request.Content, Valid: true},
	}
	postId, err := obj.dbObj.UpdatePost(c, dbPost)
	if err != nil {
		log.Printf("failed to edit post. Error: %s", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails(err.Error()))
			response.Message = "failed to edit post"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to edit post"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Post{
		PostId:   postId,
		UserId:   c.GetInt64(config.USERID),
		UserName: c.GetString(config.USERNAME),
		Content:  request.Content,
	}
	response.Message = "successfully edited post"

	c.JSON(http.StatusOK, response)
}

func (obj *postService) DeletePost(c *gin.Context) {
	var (
		request  dto.DeletePostRequest
		response dto.DeletePostResponse
	)
	if err := c.BindUri(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to delete post"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//mark post as deleted in db
	err := obj.dbObj.DeletePost(c, request.PostId, c.GetInt64(config.USERID))
	if err != nil {
		log.Printf("failed to delete post. Error: %s", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails(err.Error()))
			response.Message = "failed to delete post"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to delete post"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Post{
		PostId:   request.PostId,
		UserId:   c.GetInt64(config.USERID),
		UserName: c.GetString(config.USERNAME),
	}
	response.Message = "successfully deleted post"

	c.JSON(http.StatusOK, response)
}
