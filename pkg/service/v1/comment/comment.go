package comment

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type commentService struct {
	dbObj    db.DBLayer
	redisObj *redis.Client
}

type CommentServiceInterface interface {
	CreateComment(*gin.Context)
	ReplyComment(*gin.Context)
	GetAllReplies(*gin.Context)
	EditComment(*gin.Context)
	DeleteComment(*gin.Context)
	GetComment(*gin.Context)
	GetAllComment(*gin.Context)
}

func NewCommentService(db db.DBLayer, redisConn *redis.Client) CommentServiceInterface {
	return &commentService{
		dbObj:    db,
		redisObj: redisConn,
	}
}
