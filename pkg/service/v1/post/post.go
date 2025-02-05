package post

import (
	"liquide-assignment/pkg/db"

	"github.com/gin-gonic/gin"
)

type postService struct {
	dbObj db.DBLayer
}

type PostServiceInterface interface {
	CreatePost(*gin.Context)
	EditPost(*gin.Context)
	DeletePost(*gin.Context)
	GetPost(*gin.Context)
	GetAllPosts(*gin.Context)
}

func NewPostService(db db.DBLayer) PostServiceInterface {
	return &postService{
		dbObj: db,
	}
}
