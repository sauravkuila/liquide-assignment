package vote

import (
	"database/sql"
	"liquide-assignment/pkg/config"
	"liquide-assignment/pkg/dto"
	"log"
	"net/http"
	"time"

	e "liquide-assignment/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (obj *voteService) UpVote(c *gin.Context) {
	var (
		request  dto.UpVoteRequest
		response dto.UpVoteResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to upvote"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//TODO: add to redis and sync to db in batches

	//add the entry into db
	dbVote := dto.DbVote{
		UserId:   sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		PostId:   sql.NullInt64{Int64: request.PostId, Valid: true},
		VoteType: sql.NullBool{Bool: true, Valid: true},
	}
	voteId, err := obj.dbObj.UpsertVote(c, dbVote)
	if err != nil {
		log.Printf("failed to add vote. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to upvote"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Vote{
		VoteId:    voteId,
		PostId:    request.PostId,
		VoteType:  "upvote",
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	response.Message = "successfully upvoted"

	c.JSON(http.StatusOK, response)
}

func (obj *voteService) DownVote(c *gin.Context) {
	var (
		request  dto.DownVoteRequest
		response dto.DownVoteResponse
	)
	if err := c.BindJSON(&request); err != nil {
		log.Printf("unable to marshal request. Error:%s", err.Error())
		response.Errors = append(response.Errors, *e.ErrorInfo[e.BadRequest])
		response.Message = "failed to downvote"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//TODO: add to redis and sync to db in batches

	//add the entry into db
	dbVote := dto.DbVote{
		UserId:   sql.NullInt64{Int64: c.GetInt64(config.USERID), Valid: true},
		PostId:   sql.NullInt64{Int64: request.PostId, Valid: true},
		VoteType: sql.NullBool{Bool: false, Valid: true},
	}
	voteId, err := obj.dbObj.UpsertVote(c, dbVote)
	if err != nil {
		log.Printf("failed to add vote. Error: %s", err.Error())
		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(err.Error()))
		response.Message = "failed to downvote"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = true
	response.Data = &dto.Vote{
		VoteId:    voteId,
		PostId:    request.PostId,
		VoteType:  "downvote",
		UserId:    c.GetInt64(config.USERID),
		UserName:  c.GetString(config.USERNAME),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	response.Message = "successfully downvoted"

	c.JSON(http.StatusOK, response)
}
