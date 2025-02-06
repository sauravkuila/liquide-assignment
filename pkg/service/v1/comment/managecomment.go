package comment

import (
	"database/sql"
	"liquide-assignment/pkg/config"
	"liquide-assignment/pkg/dto"
	"log"
	"net/http"

	e "liquide-assignment/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (obj *commentService) CreateComment(c *gin.Context) {
	var (
		request  dto.CreateCommentRequest
		response dto.CreateCommentResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to add comment"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//add the entry into db
	dbComment := dto.DbComment{
		PostId:  sql.NullInt64{Int64: request.PostId, Valid: true},
		UserId:  sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		Content: sql.NullString{String: request.Content, Valid: true},
	}
	commentId, err := obj.dbObj.AddComment(c, dbComment)
	if err != nil {
		log.Printf("failed to create post. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to add comment"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Comment{
		CommentId: commentId,
		PostId:    request.PostId,
		Content:   request.Content,
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
	}
	response.Message = "successfully added comment"

	c.JSON(http.StatusOK, response)
}

func (obj *commentService) ReplyComment(c *gin.Context) {
	var (
		request  dto.CreateReplyRequest
		response dto.CreateReplyResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to add comment"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//add the entry into db
	dbComment := dto.DbComment{
		PostId:          sql.NullInt64{Int64: request.PostId, Valid: true},
		UserId:          sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		ParentCommentId: sql.NullInt64{Int64: request.ParentCommentId, Valid: true},
		Content:         sql.NullString{String: request.Content, Valid: true},
	}
	commentId, err := obj.dbObj.AddReply(c, dbComment)
	if err != nil {
		log.Printf("failed to create post. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to add comment"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Comment{
		CommentId: commentId,
		PostId:    request.PostId,
		Content:   request.Content,
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
	}
	response.Message = "successfully added comment"

	c.JSON(http.StatusOK, response)
}

func (obj *commentService) EditComment(c *gin.Context) {
	var (
		request  dto.EditCommentRequest
		response dto.EditCommentResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to edit post"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//add the entry into db
	dbComment := dto.DbComment{
		CommentId: sql.NullInt64{Int64: request.CommentId, Valid: true},
		UserId:    sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		Content:   sql.NullString{String: request.Content, Valid: true},
	}
	commentId, err := obj.dbObj.UpdateComment(c, dbComment)
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
	response.Data = &dto.Comment{
		CommentId: commentId,
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
		Content:   request.Content,
	}
	response.Message = "successfully edited post"

	c.JSON(http.StatusOK, response)
}

func (obj *commentService) DeleteComment(c *gin.Context) {
	var (
		request  dto.DeleteCommentRequest
		response dto.DeleteCommentResponse
	)
	if err := c.BindUri(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to delete comment"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//mark post as deleted in db
	err := obj.dbObj.DeleteComment(c, request.CommentId, c.GetInt64(config.USERID))
	if err != nil {
		log.Printf("failed to delete post. Error: %s", err.Error())
		if err == gorm.ErrRecordNotFound {
			response.Errors = append(response.Errors, e.ErrorInfo[e.NoDataFound].GetErrorDetails(err.Error()))
			response.Message = "failed to delete comment"
			c.JSON(http.StatusNotFound, response)
			return
		}
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to delete comment"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Comment{
		CommentId: request.CommentId,
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
	}
	response.Message = "successfully deleted comment"

	c.JSON(http.StatusOK, response)
}
