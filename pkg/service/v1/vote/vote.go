package vote

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type voteService struct {
	dbObj    db.DBLayer
	redisObj *redis.Client
}

type VoteServiceInterface interface {
	UpVote(*gin.Context)
	DownVote(*gin.Context)
}

func NewVoteService(db db.DBLayer, redisConn *redis.Client) VoteServiceInterface {
	return &voteService{
		dbObj:    db,
		redisObj: redisConn,
	}
}
