package post

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type postService struct {
	dbObj    db.DBLayer
	redisObj *redis.Client
}

type PostServiceInterface interface {
	CreatePost(*gin.Context)
	EditPost(*gin.Context)
	DeletePost(*gin.Context)
	GetPost(*gin.Context)
	GetAllPostsForUser(*gin.Context)
}

func NewPostService(db db.DBLayer, redisConn *redis.Client) PostServiceInterface {
	return &postService{
		dbObj:    db,
		redisObj: redisConn,
	}
}
