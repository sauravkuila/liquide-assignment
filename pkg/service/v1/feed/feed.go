package feed

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type feedService struct {
	dbObj    db.DBLayer
	redisObj *redis.Client
}

type FeedServiceInterface interface {
	GetFeed(*gin.Context)
}

func NewFeedService(db db.DBLayer, redisConn *redis.Client) FeedServiceInterface {
	return &feedService{
		dbObj:    db,
		redisObj: redisConn,
	}
}
