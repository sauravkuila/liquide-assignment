package comment

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
)

type commentService struct {
	dbObj db.DBLayer
}

type CommentServiceInterface interface {
	CreateComment(*gin.Context)
	EditComment(*gin.Context)
	DeleteComment(*gin.Context)
	GetComment(*gin.Context)
	GetAllComment(*gin.Context)
}

func NewCommentService(db db.DBLayer) CommentServiceInterface {
	return &commentService{
		dbObj: db,
	}
}
