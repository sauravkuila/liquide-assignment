package vote

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
)

type voteService struct {
	dbObj db.DBLayer
}

type VoteServiceInterface interface {
	UpVote(*gin.Context)
	DownVote(*gin.Context)
}

func NewVoteService(db db.DBLayer) VoteServiceInterface {
	return &voteService{
		dbObj: db,
	}
}
